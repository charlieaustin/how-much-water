var viewModel = {
    shouldShowDetails: ko.observable(false),
    name: ko.observable(),
    gallonPerCalorie: ko.observable(),
    searchFood: function(formElement) {
        var tmp = document.createElement("input");
        document.body.appendChild(tmp);
        tmp.focus();
        document.body.removeChild(tmp);
        $.get( "http://localhost:3000/data/" +formElement.food_input.value, function( data ) {
            var myData = JSON.parse(data)
            viewModel.name(myData.Name)
            viewModel.gallonPerCalorie(myData.GallonPerCalorie)
            viewModel.shouldShowDetails(true)

        });
    },
};

ko.applyBindings(viewModel);
