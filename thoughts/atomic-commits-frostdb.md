## Learning FrostDB from its commits
This is one of my ideas to learn a project. In other words, read and perhaps try interesting
commits from the project. This should help me in understanding parts of the project by focusing
on some particular fix or feature. Alternative ways e.g. getting involved with the community
or spending weeks &/ months may not be feasible unless it is part of my day job. Needless to say
this works better when the project follows atomic commits.

### Bits & Pieces
- https://github.com/polarsignals/frostdb/pull/202/commits
  - concurrency

- https://github.com/polarsignals/frostdb/pull/204/commits
  - Custom error messages + linting

- https://github.com/polarsignals/frostdb/pull/203/commits/beaeac35b004cfdf88a542985c380714da0ae475
  - from unsafe pointer to atomic pointer
  - quick insertion into btree without the need for iterator
