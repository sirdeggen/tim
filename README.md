# Lite Client Tools Demo

To demonstrate the component parts of Lite Client working together, we must spin up at minimum two peers with the full suite of tools. 
This way they may interact using all of the various protocols of communication as if acting on behalf of a merchant and customer.

## Installation

### 1.
Have [Docker](https://www.docker.com/products/docker-desktop) installed and running, and you'll also need [make](https://formulae.brew.sh/formula/make).

### 2.
```bash
git clone https://github.com/libsv/lite-client-demo.git
cd lite-client-demo
```

### 3.
Run this command to run the required containers for demonstration purposes.
```bash
make regtest
```

### 4.
Generate some blocks on regtest such that your miner has a balance to fund the customer wallet with. To do this we use a browser dashboard hosted at [http://localhost:3010](http://localhost:3010).
Use the controls to generate 101 blocks:
![regtest miner](https://bico.media/9724202bad66206d4a4e3eea4cafaa049169db220c3182331f0022ab5128d4fc.jpg)

### 5.
Run this in your terminal to fund the customer wallet:
```bash
make fund-regtest
```

### 6.
Check there is a non zero balance at the customer wallet hosted at [http://localhost:180](http://localhost:180). 


## Regtest Demo

In order to demonstrate direct payments from customer to merchant, test out the merchant example page hosted at: [http://localhost:280](http://localhost:280).
You should be able to hit the Purchase button, which generates an invoice for the Fake Cat purcahse, click the generated link to be redirected to the customer wallet such that the payment can be made. Once made the txid of the payment will appear in wallet history. Click that link to see the transaction details on the regtest explorer.

## Testnet Example

A testnet example can be found running at [https://paymail.carefulbear.com/merchant/testnet.html](https://paymail.carefulbear.com/merchant/testnet.html). Please use [https://gibsbitcoin.com](https://gibsbitcoin.com) to acquire testnet coins.

## Mainnet Example

A mainnet example can be found running at [https://paymail.carefulbear.com/merchant/mainnet.html](https://paymail.carefulbear.com/merchant/mainnet.html). Please simply send customer@carefulbear.com some BSV in order to test the mainnet version. Please do not send any more than 1000 satoshis and note that there is no way to refund coins sent to this paymail.

## Make Commands

To launch a suite of containers, simply use the make command with a single argument following:

```bash
make regtest
```

```bash
make testnet
```

```bash
make mainnet
```

To shut down the set of containers use the following command:
```bash
make stop
```

To fund the customer account on regtest:
```bash
make fund-regtest
```

## General Idea
You should be able to take the components and separate them for your own use as a place to start. For example, you could host the LiteClient folder only, and then interact with other users who are also running direct payments servers for example.

To run a mainnet Lite Client instance you'd simple copy the contents of the LiteClient folder in this repo, and run compose with the mainnet overrides, making sure to set the paymail domain to your own in the environment section of the yaml file.

```yaml
paymail:
    environment:
      DOMAIN_TLD: "your-own-domain.com" # change this to your own domain and set up your DNS records in line with the paymail standard.
```

```
docker-compose -f docker-compose.yml -f docker-compose.mainnet.yml up -d
```

There is an nginx config example file in this repo which is what we're using on the paymail.carefulbear.com domain. 
We expect you will likely host only one side of the demo components yourself (LiteClient), so you will have to set up your routing specific to your deployment.
Internal components should be able to talk to one another within the Docker bridge network.
