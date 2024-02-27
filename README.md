<h2>gRPC</h2>
o gRPC é uma poderosa ferramenta para o desenvolvimento de sistemas distribuídos, oferecendo uma maneira eficiente e flexível de comunicar-se entre serviços, independentemente das linguagens de programação ou plataformas envolvidas. Ele é amplamente utilizado em aplicações modernas, desde microserviços até grandes infraestruturas de computação em nuvem.
HTTP/2 é uma versão do protocolo HTTP que traz diversas melhorias em relação ao HTTP/1. Uma das principais mudanças é a introdução de uma conexão TCP única, que pode ser compartilhada por várias solicitações e respostas, ao contrário do HTTP/1, que estabelece uma nova conexão TCP para cada requisição e resposta. Isso resulta em uma redução significativa da sobrecarga de conexão e, consequentemente, uma melhoria no desempenho.

<h3>RPC</h3>
RPC significa Remote Procedure Call, que em português significa Chamada de Procedimento Remoto. É um modelo de programação que permite que um programa em um computador solicite serviços de um programa localizado em outro computador na mesma rede, como se estivesse chamando uma função ou método localmente.

Na prática, um RPC permite que um programa cliente execute um procedimento (ou função) em um servidor remoto, como se estivesse chamando uma função local, sem a necessidade de lidar diretamente com os detalhes da comunicação entre os sistemas. Isso simplifica o desenvolvimento de sistemas distribuídos, pois permite que os desenvolvedores se concentrem na lógica de negócios dos aplicativos, enquanto o RPC cuida da comunicação entre os diferentes componentes.

Existem diferentes implementações de RPC, incluindo o gRPC, que é uma implementação moderna e eficiente baseada no protocolo HTTP/2 e no formato de serialização de dados protobuf. Essas implementações fornecem recursos adicionais, como streaming bidirecional, autenticação e segurança, tornando o RPC uma opção poderosa para o desenvolvimento de sistemas distribuídos e comunicação entre microsserviços.

<h3>Protocolos</h3>

Uma característica importante do HTTP/2 é o conceito de Server Push, onde o servidor pode enviar múltiplas mensagens em resposta a uma única solicitação do cliente, permitindo que o cliente receba dados adicionais sem precisar solicitar explicitamente. Isso otimiza o processo de comunicação entre cliente e servidor, reduzindo a latência e melhorando a eficiência.

Além disso, o HTTP/2 utiliza a multiplexação, o que significa que várias mensagens podem ser transmitidas simultaneamente através da mesma conexão TCP. Isso permite uma utilização mais eficiente dos recursos de rede e acelera o carregamento de páginas e aplicativos web.

Em relação à estrutura dos dados, o HTTP/1 utiliza cabeçalhos de texto simples, o que pode aumentar a latência devido ao tamanho dos cabeçalhos. Já o HTTP/2 compacta os cabeçalhos e os dados em formato binário, resultando em cargas úteis mais leves e uma comunicação mais eficiente entre cliente e servidor.

Quanto ao gRPC, é um framework de comunicação remota que utiliza o HTTP/2 como seu protocolo de transporte padrão. No gRPC, o servidor opera de forma assíncrona, o que significa que pode lidar com múltiplas solicitações simultaneamente sem bloqueio. O cliente, por sua vez, pode operar de forma assíncrona ou síncrona, dependendo das necessidades da aplicação. Isso oferece flexibilidade na implementação de sistemas distribuídos e melhora o desempenho em comparação com abordagens tradicionais de comunicação remota.
