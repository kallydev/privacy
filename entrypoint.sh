#!/bin/sh

# 设置环境变量
yq w config.yaml database.tables.qq $qq | yq w - database.tables.jd $jd | yq w - database.tables.sf $sf | yq w - database.tables.wb $wb | yq w - http.host $host | yq w - http.port $port | yq w - mask $mask | sponge config.yaml

# 如果不存在数据库，创建
if [[ ! -f database/database.db ]];then
  mkdir -p database && sqlite3 database/database.db < scripts/database/create_database.sql
  if [[ -f source/6.9更新总库.txt ]];then python scripts/qq.py;fi
  if [[ -f source/www_jd_com_12g.txt ]];then python scripts/jd.py;fi
  sqlite3 database/database.db < scripts/database/create_index.sql
fi

# 如果需要重新导入qq
if [[ -f database/.reimportqq ]];then
  if [[ -f source/6.9更新总库.txt ]];then python scripts/qq.py;fi
  rm database/.reimportqq
fi

# 如果需要重新导入jd
if [[ -f database/.reimportjd ]];then
  if [[ -f source/www_jd_com_12g.txt ]];then python scripts/jd.py;fi
  rm database/.reimportjd
fi

# 启动服务器
cd server && ./app
