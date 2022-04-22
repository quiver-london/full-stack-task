# Quiver Full Stack Task

## Task


Quiver delivers orders on behalf of it's clients (merchants). An order has line items - a list of products and their respective quantity. An order can only be eligible for Quiver's delivery options if Quiver has record of the products a customer has ordered. Your task is to create a full stack application with persistent storage to support [CRUDL](https://acronyms.thefreedictionary.com/CRUDL#:~:text=Acronym,Read%2C%20Update%2C%20Delete%2C%20List) operations of products. 

### What is a product?

- A product is an named uniquely identifiable item e.g. T-Shirt, Jumper etc.
- A product is fungible: every instance of a product is equivalent. 
- A product is specific and unique to a client (a merchant working with Quiver). Products can't be shared between merchants and can't be duplicated within merchants.

### Extensions

Products are no longer entirely fungible; they can have variations/options and therefore dimensions and weights. For example a T-Shirt sold by retailer x can be purchased in any combination of sizes and colours. Your application should be updated to support this added layer of complexity.

### Technical requirements
- Front-end should be written in React. At Quiver we use Next.js, but this is by no means a requirement. 
- Back-end can be written in whatever you're most comfortable with. At Quiver we use Golang, but this is by no means a requirement.
- A database should be used, but we really don't care which type. A simple JSON db would suffice.
- Use whatever libraries you believe will achieve the best result in the shortest amount of time.

### What does a good submission look like?

- The solution is clean and simple. Less is more. There is no need to overcomplicate this task.
- REST APIs follow [Google's API style guides](https://google.aip.dev/1). CRUDL operations are detailed in AIPs [131](https://google.aip.dev/131), [132](https://google.aip.dev/132), [133](https://google.aip.dev/133), [134](https://google.aip.dev/134) and [135](https://google.aip.dev/135)
- The UI is simple, responsive and user-friendly. We like how [Stripe](https://www.youtube.com/watch?v=BwvYsLGHeRI) design their products.

## Submission Guidelines
- Fork this repository as a private repository and give [louiscollarsmith](https://github.com/louiscollarsmith) collaborator access.
- Commit your front-end and back-end code in appropriately named folders.
- Upload a walkthrough of your solution at the root of your project. We are mainly interested in the UI. At Quiver we use [Loom](https://www.loom.com/) for our product demos.

Thanks for taking the time to complete the take home task. We look forward to going through your submission!
