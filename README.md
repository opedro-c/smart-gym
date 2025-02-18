# Smart Gym
Smart Gym is a complete system designed for monitoring and recording gym equipment usage and repetitions. The system integrates hardware and software components to track exercise data in real-time.

This project utilizes an ESP32 microcontroller with multiple sensors to collect data and send it to a backend, which is built using a microservices architecture. A frontend application is also included to visualize the data in real-time.

# Architecture

### Machine / Esp32 circuit
![image](https://github.com/user-attachments/assets/5ba729d1-7f38-4fbe-b03e-3ff90ea889ed)

### Backend and ESP32 communication architecture
![image](https://github.com/user-attachments/assets/9bed1ee2-f3f3-40d5-a74b-783d84c67ed9)


### Frontend communication architecture
![image](https://github.com/user-attachments/assets/77236f1b-4bed-4a02-87af-a3600e0df5ec)




# :wrench: Technologies Used

<p>
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/cplusplus/cplusplus-original.svg" />
<img height="50px" src="https://raw.githubusercontent.com/eclipse-mosquitto/mosquitto/c85313dbde34883a150a897533d6ea5357fe3c00/logo/mosquitto-text-below.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/mongodb/mongodb-original-wordmark.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/nuxtjs/nuxtjs-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vuejs/vuejs-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/tailwindcss/tailwindcss-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-plain.svg" />
</p>
          
* Golang
* C++ / FreeRTOS
* Mosquitto / MQTT
* MongoDB
* PostgreSQL
* Nuxt + Vue.js + Tailwind
* Docker

## :rocket: Running the Project

### Machine (Esp32)
Compile and install it in a ESP32 using [Arduino IDE](https://www.arduino.cc/en/software)

### Web
Prerequisites: Docker installed

```bash
# navigate to the deploy directory
cd deploy

# Setup env variables
cp .env.example .env
export $(cat .env)

# run the application
docker compose up -d
```

# Authors
<a href="https://github.com/opedro-c">Pedro Costa Aragão</a><br>
<a href="https://github.com/talis-fb">Talison Fabio</a><br>
