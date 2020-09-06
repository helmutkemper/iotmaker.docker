package factoryDocker

import (
	"fmt"
	iotmakerDocker "github.com/helmutkemper/iotmaker.docker"
)

//  English: Small example how to use channel do monitoring image/container build
//  Português: Pequeno exemplo de como usar o canal para ver imagem/container sendo criados
func NewPullStatusMonitor() (pullStatusChannel *chan iotmakerDocker.ContainerPullStatusSendToChannel) {
	pullStatusChannel = NewImagePullStatusChannel()

	go func(c chan iotmakerDocker.ContainerPullStatusSendToChannel) {

		for {
			select {
			case status := <-c:
				// English: remove this comment to see all build status
				// Português: remova este comentário para vê todo o status da criação da imagem
				//fmt.Printf("image pull status: %+v\n", status)

				if status.Closed == true {
					fmt.Println("image pull complete!")

					// English: Eliminate this goroutine after process end
					// Português: Elimina a goroutine após o fim do processo
					return
				}
			}
		}

	}(*pullStatusChannel)

	return
}
