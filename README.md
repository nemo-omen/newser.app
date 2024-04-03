# Newser
An RSS reader with annotations. This is an in-progress application You can see a preview of what it looks like below.

# In-progress Preview
[See the (currently under development) application here](https://jellyfish-app-xbm9e.ondigitalocean.app/).

## Preview
[Quick youtube preview](https://youtu.be/zbzuPSRzj9w?si=HHutuBAeQu48H_t6)

## Very basic architecture layout

| Presenter (handlers) | Domain (usecases) | Data (repository) |
| -------------------- | ----------------- | ----------------- |
| DTOs <-------------- |                   | <-----Entities    |


Or, if you think top-down:

```
[Presentation]
      ^
      |
    (DTO)
      |
   [Domain]
      |
   (Entity)
      |
      v
 [Repository] ——> [DB]