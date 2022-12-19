# Command for running the expense api
1. docker build "docker build -t github.com/herokh/assessment:latest ."
2. docker run "docker run --name kbtg-expense-api -e PORT=:2565 -e DATABASE_URL=postgres://{user}:{pwd}@{host}/{db} -p 2565:2565 -d github.com/herokh/assessment:latest"
3. docker compose for testing sandbox "docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit --exit-code-from tests"
