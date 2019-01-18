angular.module("GruiApp").controller("questionController", [
  "$scope",
  "$state",
  "$stateParams",
  "questionService",
  "MainService",
  function questionController(
    $scope,
    $state,
    $stateParams,
    questionService,
    MainService
  ) {
    mainVm.pageName = "question";
    questionVm = this;
    questionVm.optionsCount = 4;

    questionVm.editorSetting = {
      lineNumbers: true,
      lineWrapping: true,
      indentWithTabs: true
    };

    questionVm.isCorrect = isCorrect;
    questionVm.onTagSelect = onTagSelect;

    questionVm.initCodeMirror = function() {
      $scope.cmOption = {};
      setTimeout(function() {
        $scope.cmOption = questionVm.editorSetting;
      }, 500);
    }

    questionVm.initOptionEditor = function() {
      var setting = {};
      for (var i = 0; i < questionVm.optionsCount; i++) {
        setting["option" + i] = questionVm.editorSetting;
      }
      return setting;
    }

    questionVm.getAllTags = function() {
      MainService.get("/get-all-tags").then(
        function(data) {
          if (!data || !data.data || !data.data.tags) {
            questionVm.allTags = [];
            return;
          }
          questionVm.allTags = data.data.tags;
        });
    }

    questionVm.getAllTags();

    questionVm.addNewTag = function(new_tag) {
      return {
        id: "",
        name: new_tag
      };
    }

    questionVm.validateInput = function(question) {
      if (!question.name) {
        return "Please enter valid question name";
      }
      if (!question.text) {
        return "Please enter valid question text";
      }
      if (question.positive == null || isNaN(question.positive)) {
        return "Please enter valid positve marks";
      }
      if (question.negative == null || isNaN(question.negative)) {
        return "Please enter valid negative marks";
      }
      if (Object.keys(question.options).length != questionVm.optionsCount) {
        return "Please enter all the options";
      }

      hasCorrectAnswer = false;
      correct = 0;
      angular.forEach(question.options, function(value) {
        if (value.is_correct) {
          hasCorrectAnswer = true;
          correct++;
        }
        if (!value.name) {
          return "Please enter option name correctly";
        }
      });
      if (!hasCorrectAnswer) {
        return "Please mark at least one correct answer";
      }

      if (!question.tags || !question.tags.length) {
        return "Minimum one tag is required";
      }
      if (correct > 1 && question.negative < question.positive) {
        return "For questions with multiple correct answers, negative score should be more than positive.";
      }

      return false;
    }

    function isCorrect(option, correct_options) {
      var uid = option.uid;
      if (!correct_options) {
        return false;
      }
      var optLength = correct_options.length;

      for (var i = 0; i < optLength; i++) {
        if (correct_options[i].uid == uid) {
          option.is_correct = true;
          return true;
        }
      }
      return false;
    }

    function onTagSelect(item, model) {
      for (var i = 0; i < questionVm.allTags.length; i++) {
        if (item.name == questionVm.allTags[i].name && !item.uid) {
          delete model.id;
          delete model.isTag;
          model.uid = questionVm.allTags[i].uid;
        }
      }
    }
  }
]);

angular.module("GruiApp").controller("allQuestionController", [
  "$scope",
  "$state",
  "$stateParams",
  "allQuestions",
  "questionService",
  function allQuestionController(
    $scope,
    $state,
    $stateParams,
    allQuestions,
    questionService,
  ) {
    allQVm = this;
    allQVm.searchText = "";

    allQVm.filterBy = filterBy;

    questionVm.getAllTags();

    if ($stateParams.quesID) {
      allQVm.question = allQuestions.get().find(function(q) {
        return q.uid == $stateParams.quesID;
      });
    }
    allQVm.question = allQVm.question || allQuestions.get()[0];

    allQVm.getQuestion = function getQuestion(questionId) {
      // When question is clicked on the side nav bar, we fetch its
      // information from backend and refresh it.
      questionService.getQuestion(questionId).then(function(question) {
        allQVm.question = question;
      });
    }

    allQVm.questions = function() {
      return allQuestions.get();
    }

    allQVm.toggleFilter = function(filter_value, key) {
      allQVm.filter = allQVm.filter || {};
      if (key == "tag") {
        allQVm.filter.tag = allQVm.filter.tag || [];
        var tagIndex = mainVm.indexOfObject(allQVm.filter.tag, filter_value);
        // If tag is already there in our array, then we remove it.
        if (tagIndex > -1) {
          allQVm.filter.tag.splice(tagIndex, 1);
        } else {
          allQVm.filter.tag.push(filter_value);
        }
      } else {
        allQVm.filter[filter_value] = !allQVm.filter[filter_value];
        if (filter_value == "multiple") {
          allQVm.filter.single = false;
        } else if (filter_value == "single") {
          allQVm.filter.multiple = false;
        }
      }
    }

    // TODO : Write modular code Filtering
    function filterBy(question) {
      textFilterMatch =
        question.name.toUpperCase().indexOf(allQVm.searchText.toUpperCase()) !=
        -1;

      if (allQVm.filter && allQVm.filter.tag && allQVm.filter.tag.length) {
        var found = false;
        var tagFound = true;
        for (var i = 0; i < allQVm.filter.tag.length; i++) {
          var tagIndex = mainVm.indexOfObject(
            question.tags,
            allQVm.filter.tag[i]
          );
          if (tagIndex == -1) {
            tagFound = false;
            break;
          }
          if (
            tagIndex > -1 &&
            (allQVm.filter.multiple && question.correct.length == 1)
          ) {
            tagFound = false;
          }
          if (
            tagIndex > -1 &&
            (allQVm.filter.single && question.correct.length > 1)
          ) {
            tagFound = false;
          }
          if (!tagFound) break;
        }
        return textFilterMatch && tagFound;
      } else if (allQVm.filter && allQVm.filter.multiple) {
        if (question.correct.length > 1) {
          return textFilterMatch && true;
        } else {
          return textFilterMatch && false;
        }
      } else if (allQVm.filter && allQVm.filter.single) {
        return (question.correct.length == 1) && !!textFilterMatch;
      } else {
        return !!textFilterMatch;
      }
    }

    allQVm.removeAllFilter = function removeAllFilter() {
      delete allQVm.filter;
      var questions = allQuestions.get();
      questions.length && allQVm.getQuestion(questions[0].uid);
    }
  }
]);
