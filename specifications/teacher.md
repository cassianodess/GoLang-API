# Teacher


## Get All Teachers

* Method GET

* URL: `localhost:8080/teachers`

## Get Teacher By Id

* Method GET

* URL: `localhost:8080/teachers/id`
* Ex.: `localhost:8080/teachers/1`

## Create Teacher

* Method POST

* URL: `localhost:8080/teachers`

* Payload Example: 
```
    {
        "name": "Cassiano",
        "subject": "Computer Science"
    }

```

## Update Teacher

* Method PUT

* URL: `localhost:8080/teachers/id`

* Payload Example: 
```
    {
        "name": "Julia",
        "subject": "Nursing"
    }

```

## Delete Teacher

* Method DELETE

* URL: `localhost:8080/teachers/id`
