# patients-api

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

query findPatients {
  patients {
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

query findOrders {
  orders {
    id
    message
  }
}
```

## Mutation

### createPatient

```gql
input NewPatient {
  name: String!
}

mutation createPatient {
  createPatient(input: { name: "Patient" }) {
    id
    name
    orderId
  }
}
```

### updatePatient

```gql
input UpdatePatient {
  id: ID!
  name: String
  orderId: ID
}

mutation updatePatient {
  updatePatient(input: { id: "3", name: "Patient 3", orderId: "1" }) {
    id
    name
    orderId
  }
}
```

### createOrder

```gql
input NewOrder {
  message: String!
}

mutation createOrder {
  createOrder(input: { message: "message" }) {
    id
    message
  }
}
```

### updateOrder

```gql
input UpdateOrder {
  id: ID!
  message: String!
}

mutation updateOrder {
  updateOrder(input: { id: "2", message: "new message" }) {
    id
    message
  }
}
```
