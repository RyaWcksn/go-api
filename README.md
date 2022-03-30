<h1 align="center">Interviewin API</h1>

## Introduction  

Interviewin is a platform for _Job Search_ portal for Candidates and Company can interact each other  

This is Backend service for interviewin, Created using [Golang](https://go.dev/) and for the database using [MySQL](https://www.mysql.com/)  

For the flow, bellow is the flowchart for the initial release  

```mermaid
flowchart LR;
A[User] --> B{Login}
B --> C[Company] --> F[View candidates] --> G{Accept?}
B --> D[Candidate] --> E[Apply to company] --> K{Have CV?}
G --> H[See CV and Introduction Video] --> J[Interview]
G --> I[Send email why we reject]
K --> L[Have, Upload CV]
K --> M[Not have, automate create CV]
```
