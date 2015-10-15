var viewModel = {
    searchFood: function(formElement) {
                    alert("Form submitted: " + formElement.food_input.value);
                  }
};

ko.applyBindings(viewModel);
