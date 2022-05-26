[![LinkedIn][linkedin-shield]][linkedin-url]
# Go-Bookserver
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
  </ol>
</details>

### About The Project
I am trying to practice creating backend services in Golang so I took the idea of writing a book server from [freeCodeCamp.org's example](https://github.com/AkhilSharma90/Golang-MySQL-CRUD-Bookstore-Management-API). I made major tweaks to it, albeit it started out the same.
I hope to continually add to this and extend the book server's functionality

#### Built With
<span><img height="80" src="https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/go/go.png"></span>
<span><img height="80" src="https://raw.githubusercontent.com/github/explore/80688e429a7d4ef2fca1e82350fe8e3517d3494d/topics/postgresql/postgresql.png"></span>
<span><img height="80" src="https://avatars.githubusercontent.com/u/15127678?s=280&v=4"></span>

### Description
A simple Golang server for books where you can create, retrieve, update, and delete books.

# Getting Started
1. Clone the repo
   ```sh
   git clone https://github.com/danyeric123/go-bookserver.git
   ```
2. Create a `.env` file with the following variables:
    ```
    POSTGRES_USER=postgres
    POSTGRES_PASSWORD=postgres
    POSTGRES_DB=postgres
    ```
4. Run the docker containers
    ```bash
    docker compose up
    ```
    
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/david-nagarpowers
