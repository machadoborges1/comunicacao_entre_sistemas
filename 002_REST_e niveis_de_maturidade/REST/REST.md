Documentação sobre REST

REST (Representational State Transfer) é um estilo de arquitetura para sistemas distribuídos, principalmente serviços web. Foi criado por Roy Fielding e é amplamente utilizado devido à sua simplicidade, escalabilidade e flexibilidade.  
1\. Princípios Básicos do REST  
1.1. Simplicidade

O REST adota uma abordagem simples para a comunicação entre cliente e servidor, utilizando o protocolo HTTP. A simplicidade é alcançada por meio do uso de métodos HTTP (GET, POST, PUT, DELETE) para definir as operações em recursos.

    GET: Recupera dados de um recurso.  
    POST: Envia dados para criar um novo recurso.  
    PUT: Atualiza um recurso existente.  
    DELETE: Remove um recurso.

1.2. Stateless (Sem Estado)

Um dos princípios fundamentais do REST é a comunicação stateless. Isso significa que cada requisição do cliente para o servidor deve conter todas as informações necessárias para o servidor entender e processar a solicitação. O servidor não armazena o estado da sessão do cliente entre as requisições.

**Vantagens do Stateless:**

* **Escalabilidade**: Como o servidor não precisa manter o estado de cada cliente, ele pode facilmente escalar para lidar com mais requisições.  
* **Simplificação**: Facilita a implementação, pois cada requisição é independente.

### **2\. Autenticação via Token**

Como o REST é **stateless**, a autenticação é geralmente realizada usando **tokens**. O cliente envia um token de autenticação em cada requisição para o servidor, que verifica a validade do token antes de processar a requisição.

* **JWT (JSON Web Token)**: Um formato de token amplamente utilizado que carrega informações de autenticação de forma segura e compacta.

### **3\. Conclusão**

REST é uma arquitetura simples e eficaz para serviços web, ideal para sistemas que precisam escalar e manter simplicidade. A natureza **stateless** do REST exige autenticação em cada requisição, geralmente feita via tokens como JWT, garantindo segurança e facilidade de uso.

