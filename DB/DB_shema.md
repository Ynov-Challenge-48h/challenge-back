## Client Table

| Column Name   | Data Type | Constraints |
| :------------ | :-------- | :---------- |
| client_number | INT       | PRIMARY KEY |
| client_name   | VARCHAR   |             |

## Individual Table

| Column Name             | Data Type | Constraints |
| :---------------------- | :-------- | :---------- |
| individual_id           | INT       | PRIMARY KEY |
| last_name               | VARCHAR   |             |
| first_name              | VARCHAR   |             |
| birth_date              | DATE      |             |
| cni_expiry_date         | DATE      |             |
| cni_number              | VARCHAR   |             |

## User Connection Info Table

| Column Name | Data Type | Constraints |
| :---------- | :-------- | :---------- |
| username    | VARCHAR   |             |
| uuid        | VARCHAR   |             |
| password    | VARCHAR   |             |

## Relationships

-   A Client `contains` many Individuals (One-to-Many relationship).