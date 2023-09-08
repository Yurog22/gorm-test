# Start postresql
    docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=quests -d postgres
# Requests
    GET /quests – will retrieve all quests.
    GET /quests/{id} – will fetch a quest by its id.
    POST /quest – will create a new quest.
    PUT /quests/{id} – will update a quest with the specified id.
    DELETE /quests/{id} – will delete a quest with the specified id.
# Example "POST /quest" json
    {
    "title": "Test title 1",
    "description": "Test description 1",
    "reward": 1
    }
