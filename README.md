# Go Access Log Parser

- Ler um arquivo de log a cada 5 minutos
- Guardar no mysql get/post
- Api para devolver as ultimas x entradas do método pedido
- Não pode duplicar o log(dados unicos: http_cf_connecting_ip, time_local, request, Body_bytes_send)
- Usar lib echo para http 