# Kitabisa.com

<details>
  <summary>Project Specification</summary>

  - Please make an app that shows a trivia related to the current date using the API listed here: http://numbersapi.com
  - You should make services for both the ***frontend*** and ***backend***
  - ***Backend***: get current date and use it to request the trivia from the Numbers API, return the trivia via REST API
  - ***Backend***: it should remain private, inaccessible by the public and only accessible from the ***frontend***, it should not have a public domain and/or IP
  - ***Backend***: Go is preferred, but you can use any language you want
  - ***Frontend***: request the trivia to the ***backend*** and display it
  - ***Frontend***: it should use server-side scripting (e.g. PHP, Go, Node.js, etc) because the ***backend*** should remain private
  - ***Frontend***: no preference, you can use any language and framework you want, you dont need to think about the aesthetic part of it
  - ***Frontend***: it should later be accessible through a public domain
  - Make the two services **containerized**
  - Create a **Helm chart** that can be used to deploy the containers to Kubernetes according to the rules detailed above
  - The **Helm chart** should accommodate auto-scaling but the auto-scaling should be easily toggled on and off in a specific environment, the case here is that we need auto-scaling in our production but we certainly do not in our staging and development
  - Create a **script** (hint: something as simple as a makefile will do) that will simplify the building, containerizing, and deployment in various environments (production, staging, etc)
  - Please focus more on the infrastructure side of the test, make the services (the ***frontend*** and ***backend***) as simple as you can, every single aspect that eases the deployment process will be counted
  - You DO NOT need to really deploy anything, just make sure everything works, well test it for you
  - Please put everything in a public GitHub repository and include documentation in the README
  - Be creative, but please do tell us about it in the README
</details>

## Usage Guide

## System Architecture

## Frontend
### Build Guide

## Backend
### Build Guide
### API Reference

## License
