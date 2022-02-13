# Student


## Get All Students

* Method GET

* URL: `localhost:8080/students`

## Get Student By Id

* Method GET

* URL: `localhost:8080/students/id`
* Ex.: `localhost:8080/students/1`

## Create Student

* Method POST

* URL: `localhost:8080/students`

* Payload Example: 
```
    {
        "name": "Cassiano",
        "course": "Computer Science"
    }

```

## Update Student

* Method PUT

* URL: `localhost:8080/students/id`

* Payload Example: 
```
    {
        "name": "Julia",
        "course": "Nursing"
    }

```

## Delete Student

* Method DELETE

* URL: `localhost:8080/students/id`
