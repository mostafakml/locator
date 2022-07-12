#additional questions
- What instrumentation this service would need to ensure its observability and operational
transparency?

answer:

this service run in http protocol so at the first level if drones and this service use same intranet internal ip should be accessable or at least service should be assigned 
a dns and drone also assess to that dns 
the second level each port choose for app should be exposed for drone and 
the case usind docker port should be mirror and the the port  that in actual machine should be exposed


- Why throttling is useful (if it is)? How would you implement it here?

answer:
 generally throttling is usefull for prevent over-load but for this specepic problem its seem the cost of throttling as 
 bigger than response and also the definition of problem not provide any information about 
 how mutch any drone need this service also becuase we dont implement auth
 we cant authenticate who use this service.
 so if we have those information we can add a middleware layer to implement throttling 
  
- What we have to change to make DNS be able to service several sectors at the same
time?

anwer:
at the current senario we have sector id at the config file 
if drone can have give us sectorid at the request body we can provide location for any drone in any sector

but if drone cant send us any information about sector we cant service
in other senario for example other service map for our service drone and sector id we can add other service for register
sector and drone and give access to drone the if drone is authenticated we can provide location for multiple sector
i that case we need at least 3 api service /sector  /drone  /drone-sector

we need to sore those data in a rdbms like postres and we have thre table sectors,drones
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
the strategy pattern seems a qood choice here . we can seperate those two method in our business layer then receive the method in request 
body the base run our method base on which method choose

- In general, how would you separate technical decision to deploy something from
business decision to release something?


the role of business team is what and why and the role of tech team is how
when a business decision come in tech side tech team should turn a epic to the multiple story and breaked i doable task 
priority and limitaion is also a group decision between techteam and other team that cto should 
handle them