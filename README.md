# Ask GPT

Ask GPT is a CLI tool to chat with OpenAI LLMs.

## Features

- **Custom Roles**: Customise multi roles, each with its own system prompt to tailor interactions.
- **Save/Retrieve Conversations**: Save conversations locally and continue it seamlessly in the future.
- **Streaming Output**: See the output instantly after an input, no need to wait for completion.
- **Interactive mode**: Chat with AI interactively. Offer a user experience similar to the web-based ChatGPT.
- **Support Models**: Compatibility with all models accessible through the OpenAI API.
- **Edit Config Easily**: One-click access to configuration file for easy setup and customization.

## Installation

Download the pre-built binary from [release](https://github.com/freesrz93/ask-gpt/releases/latest) and add it to path.

## Usage

1. After the first run, the app will create the config file at
    ```
    %USERPROFILE%/.config/ask-gpt/config.yaml
    ```
   for Windows and
    ```
    $HOME/.config/ask-gpt/config.yaml
    ```
   for Linux and macOS.

2. Edit the config. See [Configuration](#configuration) for detail.

3. Quickly ask a question:
    ```shell
    ag "what's the best programming language in the world?"
    ```

4. To create or continue a conversation, use `-s`:
   ```shell
   ag -s "chat 1" "Tell a story"
   ```

5. To chat interactively, use `-i`:
   ```shell
   ag -i "Translate the given sentences to Chinese"
   ```
   `-s` can be used with `-i`:
   ```shell
   ag -s "translation" -i "Translate the given sentences to Chinese"
   ```

6. To create a new role, use `-n`:
   ```shell
   ag -n
   ```
   then set its name, description and prompt interactively.

7. To use a role, use `-r` (the role should already exist):
   ```shell
   ag -r "translator" -i "Translate the given sentences to Chinese"
   ```
   `-r` can be used with `-s` but only valid when creating a conversation.

8. To show history of the conversation, use `-h`:
   ```shell
   ag -s "translation" -h -i
   ```
   this will print previous messages and continue with interactive mode.

9. For more usage, see:

   ```shell
   chatgpt --help
   ```

## Configuration

| Name              | Meaning                                                                                                    | Default                     |
|-------------------|------------------------------------------------------------------------------------------------------------|-----------------------------|
| url               | The base URL for OpenAI API.                                                                               | "https://api.openai.com/v1" |
| api_key           | Your OpenAI API key.                                                                                       | ""                          |
| model             | The model to be used.                                                                                      | "gpt-4o-mini"               |
| max_tokens        | The maximum token cost in a single API call.                                                               | 4096                        |
| temperature       | See [OpenAI Doc](https://platform.openai.com/docs/api-reference/chat/create#chat-create-temperature)       | 0.5                         |
| top_p             | See [OpenAI Doc](https://platform.openai.com/docs/api-reference/chat/create#chat-create-top_p)             | 1.0                         |
| frequency_penalty | See [OpenAI Doc](https://platform.openai.com/docs/api-reference/chat/create#chat-create-frequency_penalty) | 0                           |
| presence_penalty  | See [OpenAI Doc](https://platform.openai.com/docs/api-reference/chat/create#chat-create-presence_penalty)  | 0                           |
| editor            | The editor to be used to edit config file. Must in path.                                                   | "code"                      |
| editor_arg        | The arguments to be passed to the editor. `%path` will be replaced by the actual path.                     | "%path"                     |
