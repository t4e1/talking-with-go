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
- Conversation function flow i expected
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