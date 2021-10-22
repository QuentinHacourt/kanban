import { Component, OnInit } from '@angular/core';
import { DeveloperInput } from '../developer';
import { DeveloperService } from '../developer.service';

@Component({
  selector: 'app-register-developer',
  templateUrl: './register-developer.component.html',
  styleUrls: ['./register-developer.component.css'],
})
export class RegisterDeveloperComponent implements OnInit {
  constructor(private developerService: DeveloperService) {}

  ngOnInit(): void {}

  registerDeveloper(name: string): void {
    if (!name) {
      return;
    }
    const developerInput: DeveloperInput = {
      name: name,
    };
    this.developerService.addDeveloper(developerInput).subscribe();
  }
}
