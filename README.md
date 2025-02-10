# Smart Gym

# :wrench: Technologies Used

<p>
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" />
<img height="50px" src="https://raw.githubusercontent.com/eclipse-mosquitto/mosquitto/c85313dbde34883a150a897533d6ea5357fe3c00/logo/mosquitto-text-below.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/mongodb/mongodb-original-wordmark.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/vuejs/vuejs-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/tailwindcss/tailwindcss-original.svg" />
<img height="50px" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-plain.svg" />
</p>
          
* Golang
* Mosquitto / MQTT
* MongoDB
* PostgreSQL
* Vue.js + Tailwind
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
