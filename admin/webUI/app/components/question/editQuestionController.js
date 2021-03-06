angular.module("GruiApp").controller("editQuestionController", [
  "$state",
  "$stateParams",
  "questionService",
  "allQuestions",
  function editQuestionController(
    $state,
    $stateParams,
    questionService,
    allQuestions
  ) {
    editQuesVm = this;
    editQuesVm.newQuestion = {};

    questionVm.initCodeMirror();
    questionVm.getAllTags();

    setTimeout(function() {
      editQuesVm.editor = questionVm.initOptionEditor();
    }, 500);

    editQuesVm.markdownPreview = function() {
      return marked(editQuesVm.newQuestion.text || "");
    }

    questionService.getQuestion($stateParams.quesID).then(
      function(question) {
        var correctUids = question.correct.reduce(function(acc, val) {
          return Object.assign(acc, {[val.uid]: true});
        }, {})
        question.options.forEach(function(opt) {
          opt.is_correct = !!correctUids[opt.uid];
        })
        question.positive = parseFloat(question.positive);
        question.negative = parseFloat(question.negative);

        question.tags = question.tags || [];
        question.savedTags = question.tags.slice();

        editQuesVm.newQuestion = question;
      }
    );

    editQuesVm.cancelEdit = function() {
      if ($stateParams.returnQuizId) {
        $state.transitionTo("quiz.edit", {
          quizId: $stateParams.returnQuizId,
        });
      } else {
        $state.transitionTo("question.all", {
          quesID: editQuesVm.newQuestion.uid
        });
      }
    }

    editQuesVm.submitForm = function() {
      var question = editQuesVm.newQuestion;

      var validataionError = questionVm.validateInput(question);
      if (validataionError) {
        SNACKBAR({
          message: validataionError,
          messageType: "error"
        });
        return;
      }

      question.savedTags.forEach(function(oldTag) {
        if (!question.tags.find(function(t) { return t.uid == oldTag.uid })) {
          question.tags.push({
            uid: oldTag.uid,
            name: 'tag_to_delete uid_' + oldTag.uid,
            is_delete: true,
          })
        }
      });

      questionService.editQuestion(question).then(
        function(data) {
          allQuestions.refresh();
          SNACKBAR({
            message: data.Message
          });
          if ($stateParams.returnQuizId) {
            $state.transitionTo("quiz.edit", {
              quizId: $stateParams.returnQuizId,
            });
          } else {
            $state.transitionTo("question.all", {
              quesID: $stateParams.quesID,
            });
          }
        },
        function(err) {
          console.error(err);
          // Should not happen, but if it does remove tag deletions from the UI.
          question.tags = question.tags.filter(function(tag) {
            return !tag.is_delete;
          })
        }
      );
    }
  }
]);
