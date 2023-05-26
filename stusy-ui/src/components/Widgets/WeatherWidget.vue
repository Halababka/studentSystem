<template>
  <div id='weather' v-if="weather !== null">
    <div class="top">
      <p id='dateNowDay'>{{ weather.weather }}&deg;</p>
      <span id='wicon'>
      {{ weatherEmoji }}
    </span>

    </div>
    <p id='dateNowDay'>{{ nowDay }}, {{ weather.weatherText }}</p>

    <p>К концу занятий {{ weather.lastLessonWeather }}°,
      {{ weather.weatherStatus }}</p>

  </div>
  <div class="block skeleton" id="weather" v-else>
  </div>
</template>

<script setup>
import {computed, onMounted, ref} from "vue";
import {weatherCodes, weatherSymbols} from "@/global";

const temperaturePromise = getTemperature();
onMounted(() => {
  setTemperature(temperaturePromise);
});

const weather = ref(null);
let lastLessonTime = 1705;

const weatherEmoji = computed(() => {
  return weatherSymbols[weatherCodes[weather.value.weatherCode]];
});

const nowDay = computed(() => {
  return new Date().toLocaleString("ru", {weekday: "long"})[0].toUpperCase() + new Date().toLocaleString("ru", {weekday: "long"}).slice(1)
});

// const getDiff = computed(() => {
//   if (weather.value.weather === weather.value.lastLessonWeather) return "останется";
//   if (weather.value.weather > weather.value.lastLessonWeather) return "понизится до";
//   if (weather.value.weather < weather.value.lastLessonWeather) return "повысится до";
//   throw Error("weather check error");
// });

async function setTemperature(temperaturePromise) {
  weather.value = await temperaturePromise;
}

async function getTemperature() {
  const storedWeather = JSON.parse(localStorage.getItem("weather"));
  if (storedWeather !== null && !isWeatherExpired(storedWeather)) {
    return storedWeather;
  }
  const fetchedWeather = await fetchTemperature();
  // if (fetchedWeather[0] != "Невозможно отправить запрос") {
  //   localStorage.setItem("weather", JSON.stringify(fetchedWeather));
  //   return fetchedWeather;
  //   //   ОПАСНОСТЬ РЕКУРСИИ!
  // } else getTemperature();
  if (JSON.stringify(fetchedWeather) !== 'Невозможно отправить запрос') localStorage.setItem("weather", JSON.stringify(fetchedWeather));
  return fetchedWeather;
}

function isWeatherExpired(weather) {
  const hoursExpired = 2;
  const secondsExpired = hoursExpired * 60 * 60;
  const nowSeconds = Date.now() / 1000;
  const secondsPassed = nowSeconds - weather.fetchDate;
  return secondsPassed > secondsExpired;
}

function timeIndex() {
  let closestTime;
  let j = Math.floor(lastLessonTime / 100);
  if (j % 3 > 1) {
    closestTime = j + 1;
  } else {
    closestTime = j - j % 3;
  }
  if (closestTime === 24) {
    closestTime = 0;
  }
  return closestTime / 3;
}

function fetchTemperature() {
  const city = "Волгодонск";
  return fetch(
      `https://wttr.in/${city}?format=j1`
  )
      .then((response) => response.json())
      .then((data) => ({
        weather: parseInt(data.current_condition[0].temp_C),
        weatherText: data.current_condition[0].lang_ru[0].value.toLowerCase(),
        lastLessonWeather: parseInt(data.weather[0].hourly[timeIndex()].tempC),
        weatherCode: data.weather[0].hourly[timeIndex()]["weatherCode"],
        weatherStatus: data.weather[0].hourly[timeIndex()]["lang_ru"][0].value.toLowerCase(),
        fetchDate: Math.round(Date.now() / 1000)
      }))
      .catch((err) => {
        return ["Невозможно отправить запрос", err];
      });
}
</script>

<style scoped>
.lds-ripple {
  align-self: center;
}


#weather {
  min-height: 220px;
  flex-direction: column;
  justify-content: unset;
  gap: 30px;
  flex-wrap: nowrap;
}

#weather p, #weather h1 {
  font-size: 24px;
  margin: 0;
  z-index: 1;
}

#weather #wicon {
  /*position: absolute;*/
  /*right: -10px;*/
  /*top: -15px;*/
  font-size: 85pt;
  user-select: none;
}

#weather .top {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  height: 120px;
}

.top #dateNowDay {
  font-size: 36pt !important;
}
</style>
