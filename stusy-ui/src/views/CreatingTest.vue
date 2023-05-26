<template>
      <div class="mainContent">
        <div class="coursesMainLeft">
          <div class="redactorContentCreating-topRow">
            <h1>Название теста:</h1>
            <div class="button-save">
              <img src="@/assets/diskette.svg">
              <input type="submit" value="Сохранить"/>
            </div>
          </div>
          <label class="input-form">
            <input type="text" alt="coursesName"/>
          </label>
          <div class="redactorContentCreating-headerOfSection">
            <h1>Вопросы</h1>
            <img src="@/assets/sorting.svg"/>
            <img src="@/assets/mix.svg"/>
            <div class="button-save">
              <img src="@/assets/added.svg">
              <input type="submit" value="Создать новый"/>
            </div>
          </div>
          <div class="groupForm-row">
            <label>Текст вопроса:</label>
            <label class="input-form input-form-allWidth">
              <input type="text" alt="questionText"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Тип ответа:</label>
            <a @click="isModalOpen = true">Выбрать</a>
            <a>Повторить</a>
            <label>Количество баллов за вопрос:</label>
            <label class="input-form">
              <input type="text" alt="numberOfPoints" id="numberOfPoints"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Ответы</label>
            <img src="@/assets/addedV2.svg">
          </div>
          <div class="groupForm-row">
            <label>Ответ 1</label>
            <input type="checkbox" class="custom-checkbox" id="answer1" name="answer" value="yes">
            <label for="answer1"></label>
            <label class="input-form input-form-allWidth">
              <input type="text" alt="questionText"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Ответ 2</label>
            <input type="checkbox" class="custom-checkbox" id="answer2" name="answer" value="yes">
            <label for="answer2"></label>
            <label class="input-form input-form-allWidth">
              <input type="text" alt="questionText"/>
            </label>
          </div>
          <div class="form-group-right">
            <div class="button-save">
              <img src="@/assets/diskette.svg">
              <input type="submit" value="Сохранить"/>
            </div>
          </div>
          <ol class="accordionMenu-list">
            <li>
              <AccordionMenu class="accordionMenu" title-after="3. ZIP - это" v-bind:role="this.role"></AccordionMenu>
            </li>
            <li>
              <AccordionMenu class="accordionMenu" title-after="2. ZIP - это" v-bind:role="this.role"></AccordionMenu>
            </li>
            <li>
              <AccordionMenu class="accordionMenu" title-after="1. ZIP - это" v-bind:role="this.role"></AccordionMenu>
            </li>
          </ol>
        </div>
        <div class="coursesMainAbout">
          <h1>Настройки</h1>
          <div>
            <p>Макс.баллов:</p>
            <label class="input-form">
              <input type="text" alt="maxPoints"/>
            </label>
          </div>
          <div>
            <p>Мин. оценка в баллах:</p>
            <label class="input-form">
              <input type="text" alt="minPoints"/>
            </label>
          </div>
          <div>
            <p>Мин. оценка в баллах:</p>
            <label class="input-form">
              <input type="text" alt="minPoints"/>
            </label>
          </div>
          <div class="groupForm-row">
            <label>Равномерное распределение баллов</label>
            <input type="checkbox" class="custom-checkbox" id="UniformDistributionOfPoints"
                   name="pointsDistribution" value="yes">
            <label for="UniformDistributionOfPoints"></label>
          </div>
        </div>
      </div>
  <ModalWindow v-if="isModalOpen">
    <div class="modal-window">
      <img src="@/assets/close.svg" class="close-button" @click="isModalOpen = false">
      <h1>Выбор типа ответа</h1>
      <div class="form-group selectResponseType">
        <label class="radio-button">
          <input
              name="responseType"
              type="radio"
              value="1"
              checked
          />
          <span>Один/несколько правильных ответов</span>
        </label>
        <label class="radio-button">
          <input
              name="responseType"
              type="radio"
              value="2"
          />
          <span>Соотнести</span>
        </label>
        <label class="radio-button">
          <input
              name="responseType"
              type="radio"
              value="3"
          />
          <span>Письменный ответ</span>
        </label>
        <label class="radio-button">
          <input
              name="responseType"
              type="radio"
              value="4"
          />
          <span>Числовой ответ</span>
        </label>
        <label class="radio-button">
          <input
              name="responseType"
              type="radio"
              value="5"
          />
          <span>Расчёт по формуле</span>
        </label>
      </div>
      <textarea placeholder="Описание"></textarea>
      <div class="form-group-right">
        <div class="button-save">
          <img src="@/assets/diskette.svg">
          <input type="submit" value="Выбрать"/>
        </div>
      </div>
    </div>
  </ModalWindow>
</template>

<script>
import AccordionMenu from "@/components/AccordionMenu";
import ModalWindow from "@/components/ModalWindow";

export default {
  name: "CreatingTest",
  components: {ModalWindow, AccordionMenu},
  data() {
    return {
      role: "teacher",
      isModalOpen: false
    };
  }
};
</script>

<style scoped>
p {
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  margin-bottom: 8px;
}

h1 {
  font-style: normal;
  font-weight: 400;
  font-size: 36px;
  line-height: 41px;
}

a {
  cursor: pointer;
  font-style: normal;
  font-weight: 400;
  font-size: 20px;
  line-height: 23px;
  color: var(--blue);
}

textarea {
  background-color: var(--dark-mate-mate);
  height: 7rem;
}

.input-form, .input-form input {
  margin: 0;
}

#numberOfPoints {
  width: 40%;
  padding: 20px;
}

.custom-checkbox {
  position: absolute;
  display: block;
  z-index: -1;
  opacity: 0;
}

.custom-checkbox + label {
  cursor: pointer;
  display: block;
  align-items: center;
  user-select: none;
}

.custom-checkbox + label::before {
  content: '';
  display: inline-block;
  width: 1em;
  height: 1em;
  flex-shrink: 0;
  flex-grow: 0;
  border: 2px solid var(--blue);
  background-repeat: no-repeat;
  background-position: center center;
  background-size: 100% 100%;
}

.custom-checkbox:checked + label::before {
  display: inline-block;
  background-image: url(../assets/tick.svg);
}

.form-group-right {
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
}

.accordionMenu-list {
  gap: 20px;
}

.form-group {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 26px;
}


input[type='radio']:after {
  width: 0.8em;
  height: 0.8em;
  border-radius: 50%;
  top: -2px;
  left: -1px;
  position: relative;
  background-color: white;
  content: '';
  display: inline-block;
  visibility: visible;
  border: 2px solid var(--blue);
}

input[type='radio']:checked:after {
  width: 0.8em;
  height: 0.8em;
  border-radius: 50%;
  top: -2px;
  left: -1px;
  position: relative;
  background-color: var(--blue);
  content: '';
  display: inline-block;
  visibility: visible;
  box-shadow:  inset 0 0 0 0.15em var(--light);
}
</style>