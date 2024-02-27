# 🚀 Fastest Address API Finder

In this challenge, you'll utilize Multithreading and APIs to fetch the fastest result between two distinct APIs.

The requests will be made simultaneously to the following APIs:

- [BrasilAPI](https://brasilapi.com.br/api/cep/v1/01153000) + cep
- [ViaCEP](http://viacep.com.br/ws/" + cep + "/json/)

### Requirements:

- Respect the API that delivers the quickest response and discard the slower response.
- Display the result of the request in the command line with the address data, along with which API sent it.
- Limit the response time to 1 second. Otherwise, a timeout error should be displayed.

### Usage:

To use this program, simply run the script and enter the desired cep (Brazilian ZIP code) when prompted.