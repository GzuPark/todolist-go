# ToDo List for Go

## APIs
- [x] `GET/lists`: 목록 받기
- [x] `POST/list`: 새로운 목록 생성
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


<details>
<summary> Simple connect to database and serve </summary>

- Postgres 접속 후 `localhost:8080`에서 _Hello, World!_ 확인
- API 결과를 받기 전 단계에 status code를 받아 그대로 전달하고 custom logging 출력 함수 작성

![curl: default](./assets/curl_00.png)

</details>


<details>
<summary> Method <code>GET/lists</code> </summary>

- `GET/lists` 작성
- 현재 database에는 등록된 todo list가 없으므로 빈 목록 출력

![curl: GET/lists](./assets/curl_get_lists.png)

</details>


<details>
<summary> Apply improved error handler </summary>

- 반복해서 사용할 error handler를 범용적으로 사용할 수 있도록 변경
- panic, recover를 사용해서 error message 출력

</details>


<details>
<summary> Method <code>POST/list</code> </summary>

- `POST/list` 작성
- Request 할 때 body에 json 입력할 것

    ```json
    {
        "name": "Gopher"
    }
    ```

- 생성된 todo list의 목록을 확인할 수 있음

![curl: POST/list](./assets/curl_post_list.png)

</details>
