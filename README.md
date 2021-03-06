# ToDo List for Go

## APIs
- [x] `GET/lists`: 목록 받기
- [x] `POST/list`: 새로운 목록 생성
- [x] `GET/list/{list_id}`: 목록 이름 및 아이템 받기
- [x] `PUT/list/{list_id}`: 목록 이름 변경
- [x] `DELETE/list/{list_id}`: 목록 삭제
- [x] `POST/list/{list_id}/item`: 새로운 아이템 추가
- [x] `PUT/list/{list_id}/item/{item_id}`: 아이템 변경
- [x] `DELETE/list/{list_id}/item/{item_id}`: 아이템 삭제

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
  - `postgres://postgres:password@localhost:5432/postgres?sslmode=disable`

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


<details>
<summary> Method <code>GET/list/{list_id}</code> </summary>

- `GET/list/{list_id}` 작성
- todo list와 items를 LEFT JOIN 한 후 해당하는 list ID의 목록을 출력

![curl: GET/list/{list_id}](./assets/curl_get_list_list_id.png)

</details>


<details>
<summary> Method <code>PUT/list/{list_id}</code> </summary>

- `PUT/list/{list_id}` 작성
- 해당 list ID의 name을 변경

![curl: PUT/list/{list_id}](./assets/curl_put_list_list_id.png)

</details>


<details>
<summary> Method <code>DELETE/list/{list_id}</code> </summary>

- `DELETE/list/{list_id}` 작성
- 해당 list ID를 todo list에서 삭제

![curl: DELETE/list/{list_id}](./assets/curl_delete_list_list_id.png)

</details>


<details>
<summary> Method <code>POST/list/{list_id}/item</code> </summary>

- `POST/list/{list_id}/item` 작성
- 해당 list ID에 item 생성

    ```json
    {
        "text": "Check the progress",
        "done": false
    }
    ```

![curl: POST/list/{list_id}/item](./assets/curl_post_item.png)

</details>


<details>
<summary> Method <code>PUT/list/{list_id}/item/{item_id}</code> </summary>

- `PUT/list/{list_id}/item/{item_id}` 작성
- 해당 list ID의 특정 item 내용 수정

    ```json
    {
        "text": "Check the progress",
        "done": true
    }
    ```

![curl: PUT/list/{list_id}/item/{item_id}](./assets/curl_put_item_id.png)

</details>


<details>
<summary> Method <code>DELETE/list/{list_id}/item/{item_id}</code> </summary>

- `DELETE/list/{list_id}/item/{item_id}` 작성
- 해당 list ID의 특정 item 삭제

![curl: DELETE/list/{list_id}/item/{item_id}](./assets/curl_delete_item_id.png)

</details>
