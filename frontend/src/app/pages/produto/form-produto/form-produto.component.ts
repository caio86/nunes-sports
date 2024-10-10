import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { Produto } from 'src/app/shared/models/Produto';
import { ProdutoService } from 'src/app/shared/services/produto.service';

@Component({
  selector: 'app-form-produto',
  templateUrl: './form-produto.component.html',
  styleUrls: ['./form-produto.component.scss']
})
export class FormProdutoComponent {
  formGroup: FormGroup
  showErrors = false
  isValid = {
    id: true,
    name: true,
    description: true,
    price: true,
  }

  constructor(
    private produtoService: ProdutoService,
    private router: Router,
  ) {
    this.formGroup = new FormGroup({
      id: new FormControl("", [Validators.required, Validators.pattern("^[0-9]+$")]),
      name: new FormControl("", Validators.required),
      description: new FormControl(null),
      price: new FormControl("", Validators.required),
    })
  }

  getErrorMessage(field: string) {
    if (this.formGroup.get(field)?.hasError("required")) {
      return "Campo obrigatório"
    }

    if (this.formGroup.get(field)?.hasError("pattern")) {
      return `Valor inválido`
    }

    return "Alguma coisa deu errado"
  }

  save(): void {
    const produto: Produto = this.formGroup.value
    this.produtoService.inserir(produto).subscribe({
      next: () => {
        alert("Cliente cadastrado com sucesso!")
        this.router.navigate(["/produto"])
      },
      error: (err) => {
        console.error(err);
        alert("Erro ao cadastrar cliente!")
      }
    })
  }
}
