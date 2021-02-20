# HTTP API Specification

## /version

### GET:

#### Status codes

- 200:

    Returns the server's muwsvr version.

    Response body:

    ```json
    { "version": "string" }
    ```

## /join

### POST:

#### Request body

```json
{ "username": "username" }
```

#### Status codes

- 200:

    The user was successfully joined.

    Returns server map data and user uuid. 
    
    The map data should be rendered client-side.

    Response body:

    ```json
    { 
        "user-uuid": "71645548-7160-11eb-9439-0242ac130002",
        "chat-enabled": false,
        "map": {
            "width": 0,
            "height": 0,
            "blocks": [
                [ { "sprite-id": 0, "passable": false } ]
            ]
        } 
    }
    ```

- 400:

    Could not join the server.

    Happens when the server is already full or when the username was already taken by another user.

    Response body:

    ```json
    { "reason": "Server is full" }
    ```
    ```json
    { 
        "reason": "Username was already taken",
        "suggested-names": [ "name1", "name2", "name3" ]
    }
    ```

## /chat

### GET

#### Request body

```json
{ "user-uuid": "71645548-7160-11eb-9439-0242ac130002" }
```

#### Status codes

- 200:

    Returns chat messages.
    
    Response body:
    ```json
    { 
        "messages": [
            { "author": "username", "content": "bla bla bla" }
        ] 
    }
    ```

### PUT

#### Request body

```json
{ "user-uuid": "71645548-7160-11eb-9439-0242ac130002", "content": "bla bla bla" }
```

#### Status codes

- 200:

    Sent message.

    Returns chat messages.
    
    Response body:
    ```json
    { 
        "messages": [
            { "author": "username", "content": "bla bla bla" }
        ] 
    }
    ```

- 400:

    User is not in the world.

    Response body:
    ```json
    { 
        "reason": "User is not in the world"
    }
    ```

## /logout

### POST

#### Request body

```json
{
    "user-uuid": "71645548-7160-11eb-9439-0242ac130002"
}
```

#### Status codes

- 200:

    Successfully logged out.