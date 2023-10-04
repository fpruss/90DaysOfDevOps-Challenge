## My notes on Michael's [day03](https://www.90daysofdevops.com/2022/day03/)

I know, not new, but I think its worthy to mention, that QA testing should be automated as part of Continuous Integration.
QA is one of the big bottlenecks and it would be great to automate it. If this is possible for all applications is another question.

Its good that Michael mentions that DevOps processes and culture should be used in many positions, and not just by the  "DevOps Engineer"

## Resources

[Continuous Improvement](https://www.youtube.com/watch?v=UnjwVYAN7Ns) - Lean, Project vs Process

[Continuous Testing - IBM YouTube](https://www.youtube.com/watch?v=RYQbmjLgubM) - 
Show the upside down traditional test pyramid - Unit at the bottom, than API and on the top UI tests (so many UI tests in the traditional model).
UI changes a lot, automated UI tests will break at some point. This creates tension between "can I test it, can I change things and so on" -> anti DevOps
But automated UI tests are still way better than manual tests. The test pyramid has to be flipped, so many Unit tests, a lot of API tests and also look at
the UI tests that you had before, could we move tests from the UI level to be API tests? (will break less, will be faster and so on.
API tests also must ensure at least backwards compatibility. Then in the new model have a really low number of automated UI tests. Only real E2E validation,
do not test small feature functions on the UI level.
Still manual tests are really important - exploratory tests, UX, beta testing (i.e. roll out the change to 5% of users) 

[Continuous Integration - IBM YouTube](https://www.youtube.com/watch?v=1er2cjUq1UI) -
Integrate often to prevent merge hell. Automated build and tests.
 
[Continuous Monitoring](https://www.youtube.com/watch?v=Zu53QQuYqJ0)

[FinOps Foundation - What is FinOps](https://www.finops.org/introduction/what-is-finops/)

[The Phoenix Project: A Novel About IT, DevOps, and Helping Your Business Win](https://www.amazon.com/Phoenix-Project-DevOps-Helping-Business/dp/1942788290/)
