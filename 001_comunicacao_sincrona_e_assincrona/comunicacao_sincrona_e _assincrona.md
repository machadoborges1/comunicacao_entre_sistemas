### **1\. Comunicação Síncrona**

Na comunicação síncrona, o emissor envia uma mensagem e aguarda a resposta antes de continuar com outras tarefas. Isso significa que a execução do programa fica bloqueada até que a operação seja concluída.

### **2\. Comunicação Assíncrona**

Na comunicação assíncrona, o emissor envia uma mensagem e continua a executar outras tarefas sem esperar pela resposta imediata. A resposta é tratada de forma independente, geralmente usando goroutines e canais em Go.

### **3\. Quando Usar Comunicação Síncrona vs Assíncrona**

* **Síncrona**: Use quando o programa precisa aguardar o resultado de uma operação antes de prosseguir. É simples e fácil de entender, mas pode levar a problemas de desempenho se muitas operações bloqueantes estiverem envolvidas.  
* **Assíncrona**: Ideal para operações que podem demorar, como chamadas de rede ou processamento de dados em grande escala. Permite que outras partes do programa continuem a ser executadas enquanto aguarda o resultado.

