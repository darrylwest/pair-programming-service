# Pair Programming Service

## Overview

Create an environment whre two or more coders can work on the same files.  Private server based, Ace editor, websockets to sync files.

### Block Diagram

```mermaid
graph TD
    E1(Editor A) -- web socket --> FS((Shared File Service))
    E2(Editor B) -- web socket --> FS
    FS --> SF>SourceFile]
```

###### darryl.west | 2018.04.02

