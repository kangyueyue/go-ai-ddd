package aihelper

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/joho/godotenv"
	logger "github.com/kangyueyue/go-ai-ddd/infrastructure/common/log"

	"github.com/cloudwego/eino-ext/components/model/ollama"
	"github.com/cloudwego/eino-ext/components/model/openai"
	"github.com/cloudwego/eino/components/model"
	"github.com/cloudwego/eino/schema"
)

// StreamCallback 流式回调
type StreamCallback func(msg string)

// AIModel 模型接口
type AIModel interface {
	GenerateResponse(ctc context.Context, messages []*schema.Message) (*schema.Message, error)         // 同步生成回答
	StreamResponse(ctx context.Context, messages []*schema.Message, sc StreamCallback) (string, error) // 流式生成回答
	GetModelType() string                                                                              // 获取模型类型
}

// OpenAIModel openai模型
type OpenAIModel struct {
	llm model.ToolCallingChatModel
}

// NewOpenAIModel 创建openai模型
func NewOpenAIModel(ctx context.Context) (*OpenAIModel, error) {
	// 读取.env中的环境变量到内存中
	err := godotenv.Load()
	if err != nil {
		logger.Log.Info("Warning: .env file not found or failed to load")
		// warn：生产环境中通常不使用 .env，而是直接设置系统环境变量
	}
	key := os.Getenv("OPENAI_API_KEY")
	modelName := os.Getenv("OPENAI_MODEL_NAME")
	baseURL := os.Getenv("OPENAI_BASE_URL")

	llm, err := openai.NewChatModel(ctx, &openai.ChatModelConfig{
		APIKey:  key,
		Model:   modelName,
		BaseURL: baseURL,
	})
	if err != nil {
		return nil, fmt.Errorf("create openai model failed:%+v", err)
	}
	return &OpenAIModel{
		llm: llm,
	}, nil
}

// GenerateResponse 同步生成回答
func (o *OpenAIModel) GenerateResponse(ctx context.Context, messages []*schema.Message) (*schema.Message, error) {
	resp, err := o.llm.Generate(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("openai generate response failed:%+v", err)
	}
	return resp, nil
}

// StreamResponse 流式生成回答
func (o *OpenAIModel) StreamResponse(ctx context.Context, messages []*schema.Message, sc StreamCallback) (string, error) {
	stream, err := o.llm.Stream(ctx, messages)
	if err != nil {
		return "", fmt.Errorf("openai stream response failed:%+v", err)
	}
	defer stream.Close()

	// 流式生成流程
	var fullResp strings.Builder
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// 到达了终止符
			break
		}
		if err != nil {
			return "", fmt.Errorf("openai stream recv failed:%+v", err)
		}
		if len(msg.Content) > 0 {
			fullResp.WriteString(msg.Content) // append
			sc(msg.Content)                   // 回调
		}
	}
	return fullResp.String(), nil // 返回完整内容
}

// GetModelType 获取模型类型
func (o *OpenAIModel) GetModelType() string {
	return "openai"
}

// OllamaModel ollama模型
type OllamaModel struct {
	llm model.ToolCallingChatModel
}

// NewOllamaModel 创建ollama模型
func NewOllamaModel(ctx context.Context, baseURL, modelName string) (*OllamaModel, error) {
	llm, err := ollama.NewChatModel(ctx,
		&ollama.ChatModelConfig{
			BaseURL: baseURL,
			Model:   modelName,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("create ollama model failed:%+v", err)
	}
	return &OllamaModel{
		llm: llm,
	}, nil
}

// GetModelType 获取模型类型
func (o *OllamaModel) GetModelType() string {
	return "ollama"
}

// GenerateResponse 同步生成回答
func (o *OllamaModel) GenerateResponse(ctx context.Context, messages []*schema.Message) (*schema.Message, error) {
	resp, err := o.llm.Generate(ctx, messages)
	if err != nil {
		return nil, fmt.Errorf("ollama generate response failed:%+v", err)
	}
	return resp, nil
}

// StreamResponse 流式生成回答
func (o *OllamaModel) StreamResponse(ctx context.Context, messages []*schema.Message, sc StreamCallback) (string, error) {
	stream, err := o.llm.Stream(ctx, messages)
	if err != nil {
		return "", fmt.Errorf("ollama stream response failed:%+v", err)
	}
	defer stream.Close()

	// 流式生成流程
	var fullResp strings.Builder
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			// 到达了终止符
			break
		}
		if err != nil {
			return "", fmt.Errorf("ollama stream recv failed:%+v", err)
		}
		if len(msg.Content) > 0 {
			fullResp.WriteString(msg.Content) // append
			sc(msg.Content)                   // 回调
		}
	}
	return fullResp.String(), nil // 返回完整内容
}
