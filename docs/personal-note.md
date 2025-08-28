## Personal Notes
```Note personal thoughts for recording a development history and an after reviewing  ```

<details>
<summary>PN001. 2025-Aug-1</summary>

- Not enough resources 
    - time, skills, money...
- Use MVP based development.
- What is the MVP in my application?
    - can talk with LLM service.
- Make a **priority** of my application.
    - `1st:` Create prompt + connecting to LLM
    - `2nd:` Add tts, stt functions.  
    - `3rd:` Add another functions(user management, vocabulary note, grammar correction...). 
</details>

<details>
<summary>PN002. 2025-Aug-3</summary>

- Main function design 
- Conversation function flow which I expected
    - Client make a new conversation 
        - select conversation types(for practicing job interview, daily conversation, business conversation...)  
        - write client info(additional inforamtion for making an accurate prompting)
    - Start to talk
        - send message to API server & receive response
- At least I need two structs for binding http request.
    - One for binding user information & communcation type
    - Another for message
- The first struct should use enum type
    - return error when conversation type is diffrent from types I thought.  
</details>

<details>
<summary>PN003. 2025-Aug-7</summary>

- Created a prompt-generating function
    - The function takes the conversation type, LLM's role, user's message, using language as parameters.
- How can I preserve prompt history to maintain conversational context?
    - #1. Store the entire conversation history in memory and include it in each prompt to the LLM.
        - [ x ] Consumes too many tokens and becomes expensive.
    - #2. **Use Redis as RAG**
        - [ o ] Summarize the conversation and store the summary in Redis.
        - When generating a prompt, include the summary in the message instead of the full history.
- Should I keep the `req` field as part of the struct?
    - Not necessary for now.
        - Currently, the function only builds prompts using the client's message.
        - If additional features like logging or message transformation are added later,
then including a req field would be useful.
</details>

<details>
<summary>PN004. 2025-Aug-13</summary>

- Stream or not?
    - How to response Http Request in streaming mode? 
        - Is 1 request -> multi response possible?
        - websocket or SSE? 
    - Stream mode may not be necessary at this project
        - Purpose of application: Conversation practice by usign LLM
</details>