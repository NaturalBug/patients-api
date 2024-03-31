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
input NewPatient {
  name: String!
}

mutation createPatient($createPatientInput: NewPatient) {
  createPatient(input: $createPatientInput) {
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

mutation updatePatient($updatePatientInput: UpdatePatient) {
  updatePatient(input: $updatePatientInput) {
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

mutation createOrder($createOrderInput: NewOrder) {
  createOrder(input: $createOrderInput) {
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

mutation updateOrder($updateOrderInput: UpdateOrder) {
  updateOrder(input: $updateOrderInput) {
    id
    message
  }
}
```
