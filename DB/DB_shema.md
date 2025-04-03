## Client Table

| Column Name   | Data Type | Constraints |
| :------------ | :-------- | :---------- |
| uuid          | VARCHAR   | PRIMARY KEY |
| client_number | INT       | NOT NULL    |
| client_name   | VARCHAR   | NOT NULL    |

## Individual Table

| Column Name             | Data Type | Constraints |
| :---------------------- | :-------- | :---------- |
| uuid                    | VARCHAR   | PRIMARY KEY |
| individual_id           | INT       | NOT NULL    |
| last_name               | VARCHAR   | NOT NULL    |
| first_name              | VARCHAR   | NOT NULL    |
| birth_date              | DATE      |             |
| cni_expiry_date         | DATE      |             |
| cni_number              | VARCHAR   |             |

## User Connection Info Table

| Column Name | Data Type | Constraints |
| :---------- | :-------- | :---------- |
| uuid        | VARCHAR   | PRIMARY KEY |
| username    | VARCHAR   | NOT NULL    |
| password    | VARCHAR   | NOT NULL    |

## Relationships

-   A Client `contains` many Individuals (One-to-Many relationship).