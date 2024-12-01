
## TODO
1. Cache статистику по пользователям 

+ реализовать cache с map
`map[id]Statistic{}`

+ в cache ходим в userProvider и сохраняем в cache

+ в handler мы ходим за статистикой

+ запилить worker, который будет обновлять статистику с интервалом (30 сек)

+ если нет данных в cache то ходить в userProvider и обновлять данные