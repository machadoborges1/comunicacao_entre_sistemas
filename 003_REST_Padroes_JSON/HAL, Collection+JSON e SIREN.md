### **HAL, Collection+JSON e SIREN**

Abordaremos três formatos de hipermídia amplamente utilizados em APIs RESTful: **HAL**, **Collection+JSON** e **SIREN**. Esses formatos são projetados para adicionar estrutura e navegabilidade aos dados retornados por uma API. Também discutiremos o método HTTP `OPTIONS` e o conceito de negociações (method negotiation, content negotiation, content-type negotiation).

## **HAL (Hypertext Application Language)**

**HAL** é um formato simples que fornece uma estrutura para representar recursos e seus relacionamentos em APIs RESTful. Ele utiliza links para permitir a navegação entre recursos e facilita a descoberta dinâmica de recursos.

**Vantagens do HAL**:

* Simplicidade e facilidade de uso.  
* Facilita a navegação entre recursos de maneira padronizada.

## **Collection+JSON**

**Collection+JSON** é um formato projetado para trabalhar com coleções de recursos. Ele estrutura a resposta da API como uma coleção de itens, facilitando operações como busca, filtragem, e navegação.

**Vantagens do Collection+JSON**:

* Ideal para APIs que lidam com grandes conjuntos de dados (coleções).  
* Facilita operações em massa (ex. adição, atualização) em coleções de recursos.

## **SIREN**

**SIREN** é um formato que proporciona uma abordagem mais complexa e estruturada para representar recursos, especialmente quando os recursos possuem estados ou ações que podem ser executadas sobre eles. Ele organiza recursos em entidades que podem conter subentidades, ações, e links.

**Vantagens do SIREN**:

* Adequado para APIs complexas com recursos que têm múltiplos estados ou operações.  
* Suporte nativo para ações diretamente na resposta do recurso.

