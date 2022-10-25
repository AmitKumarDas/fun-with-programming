## Simple Designs 101

### Scenario
Take the case of a simple CLI that expects a YAML to get a job done. One may also consider this
YAML as a form of declarative API that determines the execution path.

_Note: CLI has been taken for simplicity purposes. This upholds for a long-running service as well._

### Typical Design Process
Developers will go great length to brainstorm on the schema for this YAML. Canonical approach is
to strive for a schema that is generic enough to accept new features. Let us take an example to
clarify this further.

### Sample YAML
Below is a sample YAML that is understood by our CLI. In this case CLI can be provided either
with an image or an image-bundle. In addition, this artifact (i.e. image or bundle) can be found
in some file path or some registry. Finally, the found artifact should be saved to some destination
which in turn can either be a file location or some registry location. Note that ?future implies
the extensibility of this schema. In other words, in future one can add more types or location or
both & the schema will remain same. Sleek!
```yaml
kind: MyCLI
version: v1.0.0
spec:
  type: image | image-bundle | ?future
  source: file | registry | ?future
  dest: file | registry | ?future
```

### Typical Implementation
Let us think about possible logic to implement our well-designed schema. Needless to say this logic
must take care of type x source x dest combinations. In other words, there will be multiple if else
conditions. The careful reader might point out that there will multiple nested if else conditions.
If we consider using abstractions, then we might implement interfaces e.g. `Location` with concrete
implementations like file & registry. This might further push us writing some switch case 
(read conditional) logic to switch between these concrete implementations based on values set in the
YAML.

In the future, when new categories arrive (i.e. new type or source or dest), we need to change existing
implementations to accommodate these new arrivals.

### Are we good?
Well I would like to analyse the thought process that was very common across above sections. In my
humble opinion, we gave a lot of stress to having a single schema for every purpose. Instead, I would
worry more about the principles that we violated. For example, most of the above logic does not adhere
to Single Responsibility Pattern nor is it better in regard to Open Closed principle. This design seems
biased towards LID parts of SOLID principles.

### What if we focus just on Single Responsibility & Open Closed Principles
Let us go back to our schema design. How about accommodating multiple schema each with a specific 
purpose. Let us review the following YAMLs which is a result of this new thought process.

```yaml
kind: ImgFileToReg
version: v1.0.0
spec:
  source: file-path-to-source-file # can only be a file path
  dest: registry-path-to-destination # can only be a registry path
```

```yaml
kind: ImgRegToFile
version: v1.0.0
spec:
  source: registry-path-to-source # can only be a registry path
  dest: file-path-to-destination # can only be a file path
```

```yaml
kind: BundleFileToReg
version: v1.0.0
spec:
  source: file-path-to-source # can only be a file path 
  dest: registry-path-to-destination # can only be a registry path
```

```yaml
# Let this file be named as abc.yaml
kind: BundleFileToRegList
version: v1.0.0
spec:
  - source: abc-1.tar # can only be a file path
    dest: pkg/abc-1 # can only be a registry path
  - source: abc-2.tar # can only be a file path
    dest: pkg/abc-2 # can only be a registry path
```

In above YAMLs, we find that each yaml follows Single Responsibility as well as Open Closed principles.
In addition, we are not bothered about the future. In other words, we create a new schema only when there
is a real need. 

However, it might make sense to have a mapper YAML to let our CLI decide the respective
function to be invoked given a YAML. A mapper YAML might look like the following:

```yaml
kind: FileToKindMapper
version: v1.0.0
spec:
  abc.yaml: BundleFileToReg/v1.0.0
  def.yaml: ImageRegToFile/v1.0.0
  abc-2.yaml: BundleFileToReg/v2.0.0
```

### A YAML corresponding to an endpoint
Above separation of concerns right from the schema will help us to avoid nested conditionals
& also negate the need for abstractions (read interfaces & associated logic). One may argue that the 
developer experience might not be great. After all user needs to deal with variety of YAMLs to get their
job done. On the other side, we could counter that in case of older schema this user must understand every
field & set them appropriately to get the job done. On the whole even though DevEx is crucial, it can not
be the deciding factor for our implementation.

Note that YAML & 12 factors are a rage these days. There is nothing bad in them. However, it should not
come at the cost of diluting SOLID programming principles. In this particular scenario a YAML is the 
declarative API. However, an API design does not cram all the methods together & instead provides an
endpoint for a specific purpose. Can we think on similar lines for declarative APIs as well? In other
words, one YAML defined to execute a specific responsibility.

### Concluding remarks
Let us pause thinking about software development & start looking from end users perspective. What do the users 
need? What kind of software is loved by an operator? Let us try writing down their train of thoughts.

Will it be possible to support old as well new versions of schema while still using the latest binary? Will it be
possible to cut new releases fast? Is it easy to add new test cases to older releases & hence aid in reproducing
customer issues? Can the customer upgrade from a very old binary, e.g. latest - 5 to the latest version seamlessly?
Is it feasible for the customer to run multiple versions of the binary where one version is for the existing
use cases while a newer version is for an experimental use case. In other words, can the user do a phased rollout
of the binary in their production environment. The list continues. These real world scenarios can be endless. 

The approach that can meet most of them is to have the logic minus cyclomatic complexity, minus abstractions, plus
modular logic, plus good practices, plus simple patterns that help in marching towards the goal, minus the patterns
that don't fit current context.

Needless to say I will bet my money on one YAML one purpose approach which I believe can tackle the maintainability
aspects better.
