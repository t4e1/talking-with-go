// readme
# Speaking Helper
## Outline
- `purpose:` A personal project using Go + C + LLM services.
- `simple architecture:` Simple Client + Web Application + Monitoring 
- `skills:` Go, C(eBPF), AWS(or other cloud services?), promstack, postgresql, and so on...
- `theme:` When the speaker choose specific topic of conversation and talks about it, LLM returns appropriate answer to the Speaker. 
- `predicted Workflow:` Speaker > (voice to text) > ollama(or ?) > (text to voice) > Speaker 

## Design Point
<details>
<summary>DP001. 2025-Aug-1</summary>

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

## Tech stack 
- goloang 1.24.5 (latest)
- gin
    - gin vs. echo vs. chi 
    - good to create light applicaton.
    - a wealth of learning materials.

## History
- `2025.08.01` Project start

