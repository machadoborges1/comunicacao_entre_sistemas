### **Documentação sobre Níveis de Maturidade REST**

A Maturidade REST é um conceito desenvolvido por Leonard Richardson, que define quatro níveis de conformidade para APIs RESTful, chamados de Níveis de Maturidade REST. Cada nível representa um grau crescente de conformidade com os princípios REST, culminando em uma API completamente RESTful.

#### **Nível 0: The Swamp of POX**

No Nível 0, uma API usa HTTP apenas como um meio de transporte, mas não adota nenhuma das características do protocolo, como métodos HTTP, URIs significativos, ou estrutura de resposta padrão. Este nível é frequentemente chamado de **"The Swamp of POX"** (Plain Old XML), onde as APIs se comportam como chamadas de procedimento remoto (RPC).

#### **Nível 1: Recursos**

No Nível 1, a API começa a usar recursos e URIs (Uniform Resource Identifiers) para identificar as entidades manipuladas, mas ainda não utiliza adequadamente os métodos HTTP. Cada recurso tem seu próprio endpoint.

**Exemplo:**

* `GET /books/123`: Retorna o livro com ID 123\.  
* `POST /books`: Cria um novo livro.

Este nível ainda pode ser visto como RPC, mas com um design de URI mais claro.

#### **Nível 2: Verbos HTTP**

No Nível 2, a API adota totalmente os métodos HTTP (GET, POST, PUT, DELETE), associando-os a diferentes operações de CRUD (Create, Read, Update, Delete) em recursos.

**Exemplo:**

* `GET /books/123`: Recupera o livro com ID 123\.  
* `POST /books`: Cria um novo livro.  
* `PUT /books/123`: Atualiza o livro com ID 123\.  
* `DELETE /books/123`: Deleta o livro com ID 123\.

#### **Nível 3: HATEOAS (Hypermedia As The Engine Of Application State)**

No Nível 3, a API não só usa recursos e verbos HTTP corretamente, mas também adota HATEOAS. Isso significa que as respostas incluem links para outras operações possíveis, permitindo que o cliente navegue dinamicamente pela API.

**Exemplo:**

* A resposta de `GET /books/123` pode incluir links para atualizar ou deletar o livro, como:

### **Conclusão**

Os Níveis de Maturidade REST de Richardson mostram a evolução de uma API simples e sem recursos para uma API completamente RESTful, que adota plenamente os princípios REST e facilita a interação com os clientes de forma padronizada e eficiente.

1. **Nível 0 (Swamp of POX)**: Simplesmente utiliza HTTP como transporte.  
2. **Nível 1 (Recursos)**: Introduz URIs que representam recursos.  
3. **Nível 2 (Verbos HTTP)**: Usa métodos HTTP corretamente para interagir com os recursos.  
4. **Nível 3 (HATEOAS)**: Fornece links de navegação para os clientes, promovendo a descoberta dinâmica de operações.

Cada nível tem seu propósito e pode ser adequado dependendo do contexto e das necessidades do sistema que está sendo desenvolvido.
