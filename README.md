# go.message-tagger
<a href="https://www.repostatus.org/#wip"><img src="https://www.repostatus.org/badges/latest/wip.svg" alt="Project Status: WIP – Initial development is in progress, but there has not yet been a stable, usable release suitable for the public." /></a><br>

This service is a part of the repo: [microservices-training-ground](https://github.com/JacekKorta/microservices-training-ground)<br>
This service reads messages (stackoverflow questions) from the queue, checks the tags and looks for phrases in the question body. Depending on the results, the service adds its own tags (in the “reasons” list) and requeues enrichment messages. 

### How to run?

You should run this service via docker compose in main repo [microservices-training-ground](https://github.com/JacekKorta/microservices-training-ground)

Create env file:
```bash
cp .env.example .env
```

Additional info about env variables:

BARE_TAG - eg. “laravel” - service will mark the messages if they don't have the bare tag. 
Eg. if you subscribe “laravel-8” tag via go-stack-questions and someone asks a question with the tag “laravel-8” then you will receive an info if this message has no bare_tag:”laravel”<br>
DESIRABLE_TAG - you will receive info for each question with this tag.<br>
WARNING_STRINGS - you will receive info for each question with the string configured here. If you want to receive messages for more than one string then you must  put more phrases and split them by pipe “|” eg.: “first string|second string|third string” 
