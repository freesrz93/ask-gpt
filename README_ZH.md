# Ask GPT

[English](./README.md) | 简体中文

Ask GPT 是一个命令行工具, 用于通过 API 与 OpenAI 的 GPT 模型聊天.

## 功能特点

- **自定义角色**: 可自定义多个角色, 每个角色都有自己的 prompt.
- **保存/检索对话**: 将对话保存到本地, 并在稍后无缝继续.
- **流式输出**: 输入后即时看到输出, 无需等待完成.
- **管道支持**: 通过管道将任何其他命令行工具的输出发送给 AI.
- **交互模式**: 与 AI 交互聊天, 提供类似于基于网页的 ChatGPT 的用户体验.
- **多种模型**: 兼容通过 OpenAI API 访问的所有模型.
- **编辑文件配置**: 一键编辑配置文件, 便于设置和自定义.

## 安装

从 [release](https://github.com/freesrz93/ask-gpt/releases/latest) 下载对应平台的预构建二进制文件并添加到 path 中.

## 使用

1. 首次运行 ask-gpt:
   ```shell
   ag -v
   ```
   将在以下位置创建配置文件  
   Windows:
    ```
    %USERPROFILE%/.config/ask-gpt/config.yaml
    ```
   Linux/macOS:
    ```
    $HOME/.config/ask-gpt/config.yaml
    ```

2. 编辑配置文件. 有关详细信息, 参见 [配置](#配置).

3. 快速提问:
    ```shell
    ag "世界上最好的编程语言是什么？"
    ```

4. 使用 `-s` 创建新的/继续已有对话:
   ```shell
   ag -s "chat 1" "讲个故事"
   ```

5. 使用 `-i` 互动聊天:
   ```shell
   ag -i "将以下句子翻译为中文"
   ```
   `-s` 可以与 `-i` 一起使用:
   ```shell
   ag -s "translation" -i "将以下句子翻译为中文"
   ```

6. 使用管道:
   ```shell
   cat raw.txt | ag -i "将上述内容翻译为中文"
   ```

7. 使用 `-n` 创建新角色:
   ```shell
   ag -n
   ```
   然后可以交互式的设置其名称、描述和提示.

8. 使用 `-r` 使用角色（角色需存在）:
   ```shell
   ag -r "translator" -i "将以下句子翻译为中文"
   ```
   `-r` 可以与 `-s` 一起使用, 但仅在创建对话时有效.

9. 使用 `-h` 显示对话历史:
   ```shell
   ag -s "translation" -h -i
   ```
   这将打印之前的消息并进入交互模式.

10. 更多用法参见:
    ```shell
    chatgpt --help
    ```

## 配置

| 名称                | 意义                                                                                                      | 默认值                         |
|-------------------|---------------------------------------------------------------------------------------------------------|-----------------------------|
| url               | OpenAI API 的 URL.                                                                                       | "https://api.openai.com/v1" |
| api_key           | OpenAI API 密钥.                                                                                          | ""                          |
| model             | 要使用的模型.                                                                                                 | "gpt-4o-mini"               |
| max_tokens        | 单次 API 调用消耗的令牌数上限.                                                                                      | 4096                        |
| temperature       | 见 [OpenAI 文档](https://platform.openai.com/docs/api-reference/chat/create#chat-create-temperature)       | 0.5                         |
| top_p             | 见 [OpenAI 文档](https://platform.openai.com/docs/api-reference/chat/create#chat-create-top_p)             | 1.0                         |
| frequency_penalty | 见 [OpenAI 文档](https://platform.openai.com/docs/api-reference/chat/create#chat-create-frequency_penalty) | 0                           |
| presence_penalty  | 见 [OpenAI 文档](https://platform.openai.com/docs/api-reference/chat/create#chat-create-presence_penalty)  | 0                           |
| editor            | 用于编辑配置文件的编辑器. 必须在 path 中.                                                                               | "code"                      |
| editor_arg        | 传递给编辑器的参数. `%path` 将被替换为实际配置文件路径.                                                                       | "%path"                     |
