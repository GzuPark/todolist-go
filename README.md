# ToDo List for Go

## APIs
- [ ] `GET/lists`: 목록 받기
- [ ] `POST/list`: 새로운 목록 생성
- [ ] `GET/list/{list_id}`: 목록 이름 및 아이템 받기
- [ ] `PUT/list/{list_id}`: 목록 이름 변경
- [ ] `DELETE/list/{list_id}`: 목록 삭제
- [ ] `POST/list/{list_id}/item`: 새로운 아이템 추가
- [ ] `PUT/list/{list_id}/item/{item_id}`: 아이템 변경
- [ ] `DELETE/list/{list_id}/item/{item_id}`: 아이템 삭제

## Developing steps

<details>
<summary> Prepare the database server </summary>

- Software
  - Docker
  - Insomnia or Postman
- Pull `postgres` docker images

    ```sh
    docker pull postgres
    ```
- Create `tododb` container

    ```sh
    docker run -d --name tododb \
        -e POSTGRES_PASSWORD=password \
        -p 5432:5432 \
        postgres
    ```

- Create SQL tables

    ```sh
    docker cp ./assets/schema.sql tododb:/docker-entrypoint-initdb.d/schema.sql
    docker exec -u postgres tododb psql postgres -U postgres -f docker-entrypoint-initdb.d/schema.sql
    ```

- Connection URL
  - `postgres://postgres:password@localhost:5432/postgres?sslmode-disable`

</details>
