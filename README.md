
## TODO
1. Добавить created, updated, deleted у chat 

добавить папку ./migrations -> `<YYYYMMDDHHMMSS>_<название миграции>.sql`
(`alter table chat add column ...`)
со стороны кода проверить всё

2. Cache статистику по пользователям 

добавить ручку GET ./user/stats/:id
добавить ручку GET ./user/stats

+ реализовать cache с map
`map[id]Statistic{}`

+ реaлизация похода за статистикой в userProvider

+ в cache ходим в userProvider и сохраняем в cache

+ в handler мы ходим за статистикой

```
cacheProvider interface = userProvider interface

handler(userProvider)
handler(cacheProvider(userProvider))

handler - cache (decorator) - userProvider
```

+ запилить worker, который будет обновлять статистику с интервалом (30 сек)

+ если нет данных в cache то ходить в userProvider и обновлять данные

+ понадобятся mutex (rw mutex)