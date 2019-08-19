---
title: Append Arrays
weight: 4609
---

# Append Arrays
This activity allows you to concatenate two arrays

## Installation
### Flogo Web
Zip the parent folder (appendarraysv2)  and upload it via the UI
### Flogo CLI
```bash
flogo add activity github.com/TIBCOSoftware/flogo-contrib/activity/counter
```

## Schema
Inputs and Outputs:
```json
{
  "input":[
    {
      "name": "array1",
      "type": "array",
      "required": true
    },
    {
      "name": "array2",
      "type": "array",
      "required": true
    }
  ],
  "output": [
    {
      "name": "output",
      "type": "array",
      "required": true
    }
  ]
}
```
