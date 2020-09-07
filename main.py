from datetime import datetime

print(datetime.now().strftime('%m{y}%H{h}%M{m}').format(y='月', h='日', m='分'))