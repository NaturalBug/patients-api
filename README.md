# patients-api

The project is built for
https://github.com/NaturalBug/patients-management-system

## Todo

- Implement functions in postgres.go
- schema.resolvers.go use `postgres` to get data in DB and get rid of using patients and orders in Resolver

## Usage

### Run at local

```sh
make run
```

### Generate code after schema changed

```sh
make generate
```

## Query

### Patients

```gql
type Patient {
  id: ID!
  name: String!
  orderId: ID
}

query findPatients($id: ID) {
  patients(id: $id) {
    id
    name
    orderId
  }
}
```

### Orders

```gql
type Order {
  id: ID!
  message: String
}

query findOrders($id: ID) {
  orders(id: $id) {
    id
    message
  }
}
```

## Mutation

### createPatient

```gql
input CreatePatientInput {
  name: String!
}

mutation createPatient($createPatientInput: CreatePatientInput) {
  createPatient(input: $createPatientInput) {
    id
    name
    orderId
  }
}
```

### updatePatient

```gql
input UpdatePatientInput {
  id: ID!
  name: String
  orderId: ID
}

mutation updatePatient($updatePatientInput: UpdatePatientInput) {
  updatePatient(input: $updatePatientInput) {
    id
    name
    orderId
  }
}
```

### createOrder

```gql
input CreateOrderInput {
  message: String!
}

mutation createOrder($createOrderInput: CreateOrderInput) {
  createOrder(input: $createOrderInput) {
    id
    message
  }
}
```

### updateOrder

```gql
input UpdateOrderInput {
  id: ID!
  message: String!
}

mutation updateOrder($updateOrderInput: UpdateOrderInput) {
  updateOrder(input: $updateOrderInput) {
    id
    message
  }
}
```

## PostgreSQL

### patients

```sql
create table patients
(
	Id serial not null,
	Name varchar(100),
	OrderId int,
	primary key(Id),
	constraint fk_OrderId
		foreign key(OrderId)
		references orders(Id)
)
```

### orders

```sql
create table orders
(
	Id serial not null,
	Message Text
)
```
