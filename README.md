# Paytm integration using Golang 

Read about how we made this on our blog (https://www.mindinventory.com/blog/paytm-integration-using-golang/)

<img src="https://raw.githubusercontent.com/Mindinventory/Golang-Paytm/master/paytm-new.png">

## Requirements : 
 
* i.  Provide the value for PAYTM_MERCHANT_KEY in .env file. (The value for MERCHANT_KEY will be provided after the onboarding process is completed).
     
## Installation steps:

* i. For successful transaction, Paytm generates checksumhash and verifies this checksumhash on  merchant server.
     
     Copy provided library.go file to your local. which is used for the Checksum generaion and verification. 

  ii. Paytm integration can be done in two stages (staging and production modes).
      staging and production modes have different Paytm transaction urls.
  
  iii. In staging mode we are using sandbox details for testing the paytm working.
       
       Staging mode transaction url 
       https://securegw-stage.paytm.in/theia/processTransaction 
       
   iv. If you are using production mode means it accepts original details of paytm. 
  
        Production mode transaction url    
        https://securegw.paytm.in/theia/processTransaction
        
  
  For more details about Paytm refer this link 
  https://developer.paytm.com/docs/v1/payment-gateway  
     
  
## LICENSE!

PayTm integration using Golang is [MIT-licensed](https://github.com/mindinventory/Golang-Paytm/blob/master/LICENSE)

## Let us know!
We’d be really happy if you sent us links to your projects where you use our component. Just send an email to sales@mindinventory.com And do let us know if you have any questions or suggestion regarding our work.



  
    
