#additional questions
- What instrumentation this service would need to ensure its observability and operational
transparency?

answer:

this service run in http protocol so at the first level 
if drones and this service use same intranet network internal ip should be accessable 
or at least service should be assigned 
to dns domain like (data-bank.com) and drone also assess to that dns 
the second level each port choose for app should be exposed for drone and 
the case using docker port should be mirror and  the that port   in actual machine should be exposed


- Why throttling is useful (if it is)? How would you implement it here?

answer:
 generally throttling is useful for prevent over-load but for this 
 specific problem its seem the cost of throttling is 
 bigger than response  also the definition of problem not provide any information about 
 how much any drone need this service also because we dont implement auth
 we cant authenticate who use this service.
 so if we have this information we can add a middleware layer to implement throttling 
 other option is implement throttling base on ip how many a ip can use our service  this case is also 
 doable by adding a custom middleware
  
- What we have to change to make DNS be able to service several sectors at the same
time?

answer:
at the current scenario we have sector id at the config file 
if drone can have given us sectorid at the request body we can provide location for any drone in any sector

but if drone cant send us any information about sector we cant service
in other scenario for example third service map for our service drone and sectors
 we can add other api service for register
sector and drone and give access to drone the if authenticated drone we can provide location for multiple sector
i that case we need at least 3 api service /sector  /drone  /drone-sector

we need to sore those data in a rdbms like postreSQL and we have two table sectors,drones
and the relation between drones and sector is one to many 
also we need jwt service to authenticate our drones

- Our CEO wants to establish B2B integration with Mom's Friendly Robot Company by
allowing cargo ships of MomCorp to use DNS. The only issue is - MomCorp software
expects loc value in location field, but math stays the same. How would you
approach this? What’s would be your implementation strategy?

answer:
we can add a decorator layer to change our response in multiple way and serve 

- Atlas Corp mathematicians made another breakthrough and now our navigation math is
even better and more accurate, so we started producing a new drone model, based on
new math. How would you enable scenario where DNS can serve both types of clients?

answer: 
the strategy pattern seems a qood choice here . we can seperate those two method in our business layer
 then receive the method in request body as property
  base on property our service decide use which method

- In general, how would you separate technical decision to deploy something from
business decision to release something?


the role of business team is what and why and the role of tech team is how
when a business decision come in tech side tech team should turn a epic to the multiple story and breaked i doable task 
priority and limitation is also a group decision between tech team and other team that cto should 
handle them