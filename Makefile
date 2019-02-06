# GET
accounts_filter=curl http://127.0.0.1:8080/accounts/filter
accounts_group=curl http://127.0.0.1:8080/accounts/group
accounts_1_recommend=curl http://127.0.0.1:8080/accounts/1/recommend
accounts_1_suggest=curl http://127.0.0.1:8080/accounts/1/suggest

# POST
# accounts_new=curl -d "birthyear=1905&press=%20OK%20" http://127.0.0.1:8080/accounts/new
accounts_new=curl -d '{"sname": "Хопетачан","country": "Голция","birth": 736598811,"id": 50000,"email": "orhograanenor@yahoo.com","sex":"f","fname": "Полина"}' -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/accounts/new

# accounts_new=curl -d '\
# {\
#     "sname": "Хопетачан",\
#     "email": "orhograanenor@yahoo.com",\
#     "country": "Голция",\
#     "interests": [],\
#     "birth": 736598811,\
#     "id": 50000,\
#     "sex": "f",\
#     "likes": [\
#         {"ts": 1475619112, "id": 38753},\
#         {"ts": 1464366718, "id": 14893},\
#         {"ts": 1510257477, "id": 37967},\
#         {"ts": 1431722263, "id": 38933}\
#     ],\
#     "premium": {"start": 1519661251, "finish": 1522253251},\
#     "status": "всё сложно",\
#     "fname": "Полина",\
#     "joined": 1466035200\
# }\
# ' -H "Content-Type: application/json" -X POST http://127.0.0.1:8080/accounts/new
accounts_likes=curl -d "birthyear=1905&press=%20OK%20" http://127.0.0.1:8080/accounts/likes
accounts_1=curl -d "birthyear=1905&press=%20OK%20" http://127.0.0.1:8080/accounts/1

get:
	@ $(accounts_filter)
	@ $(accounts_group)
	@ $(accounts_1_recommend)

post:
	@ $(accounts_new)
	@ $(accounts_likes)

all:
	@ $(accounts_filter)
	@ $(accounts_group)
	@ $(accounts_1_recommend)
	@ $(accounts_1_suggest)
	@ echo ""
	@ $(accounts_1)
	@ $(accounts_new)
	@ $(accounts_likes)
