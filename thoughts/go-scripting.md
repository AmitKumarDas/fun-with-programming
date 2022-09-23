## Motivation
I have often found people talking of replacing their bash scripts with a higher level language 
that offers predictability. I had never encountered such a need until I faced a situation
where the number of lines in bash script exceeded 600. 

### Take 1 (Aug 2022)
During these days I had been watching a golang project called testscript. I thought this go based 
project can help me realise the goodness of bash along with correctness of go. I started 
experimenting & soon came ot realise that testscript is not the right tool to replace bash. As of
the latest news, it might be a good tool to test CLIs.

### Take 2 (Aug 2022)
Meanwhile, I was wondering if task based tool will be a good replacement for bash. I knew mage is 
a go based tool that can be used instead of make. However, somewhere in my mind I knew that task
based executions do not really solve the scripting needs. It might backfire in the worst cases.

### Take 3 (Sep 2022)
I started digging more into the internals of mage & found its sh package to be quite useful. It is
not a tool to replace bash but provides just enough functions to make scripting in go effective.
I started making few changes to mage's sh package to make it aligned for scripting business. 
For example, I wanted the logic to detect missing ENV variables which are so common in scripts.
This pushed me to add several utility functions to automagically expand ENV variables if present &
hence make the whole scripting journey smooth. In addition, I never had to bother about linters, etc.
since go compiler was more than enough.

### Take 4 (Sep 2022)
How about using `go test` as the UX to trigger the newly built go script?