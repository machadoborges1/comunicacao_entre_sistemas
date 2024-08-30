## **HTTP `OPTIONS`**

O método HTTP `OPTIONS` é usado para descrever as opções de comunicação disponíveis para um recurso especificado. Quando um cliente faz uma requisição `OPTIONS` para um endpoint, o servidor retorna os métodos HTTP suportados para aquele endpoint, junto com outras informações, como cabeçalhos permitidos.

Permite que o cliente descubra as capacidades do servidor antes de tentar uma operação.  
Útil para navegadores na implementação de CORS (Cross-Origin Resource Sharing).

## **Method Negotiation, Content Negotiation e Content-Type Negotiation**

**Method Negotiation** refere-se à escolha do método HTTP apropriado para interagir com um recurso, como `GET`, `POST`, `PUT`, ou `DELETE`. Essa escolha é feita pelo cliente com base nas operações que deseja realizar.

**Content Negotiation** é o processo pelo qual o cliente e o servidor concordam sobre o formato dos dados que serão trocados. Isso é feito através do cabeçalho `Accept` na requisição, onde o cliente indica quais formatos ele pode aceitar (por exemplo, `application/json`, `application/xml`). O servidor então responde no formato que melhor corresponde ao pedido do cliente.

**Content-Type Negotiation** refere-se ao servidor determinando o formato correto de resposta com base no cabeçalho `Accept` fornecido pelo cliente. Isso permite que o servidor ofereça a resposta no formato mais adequado para o cliente, seja JSON, XML, ou outro.

### **Conclusão**

Os formatos HAL, Collection+JSON, e SIREN são diferentes abordagens para implementar hipermídia em APIs RESTful, cada um com suas próprias vantagens e casos de uso. Junto com a compreensão do método `OPTIONS` e das diferentes formas de negociação (method negotiation, content negotiation, e content-type negotiation), essas ferramentas permitem que APIs sejam mais robustas, flexíveis e autodocumentadas.

